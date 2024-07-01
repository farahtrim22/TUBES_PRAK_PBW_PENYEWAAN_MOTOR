package motorcontroller

import (
	"Tubes_PBW/config"
	"Tubes_PBW/entities"
	"Tubes_PBW/libraries"
	"Tubes_PBW/models"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"
)

var validation = libraries.NewValidation()
var MotorModel = models.NewMotorModel()
var customerModel = models.NewCustomerModel()

func Index(response http.ResponseWriter, request *http.Request) {

	session, _ := config.Store.Get(request, config.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(response, request, "/login", http.StatusSeeOther)
	} else {
		if session.Values["loggedIn"] != true {
			http.Redirect(response, request, "/login", http.StatusSeeOther)
		} else {
			motor, _ := MotorModel.FindAll()
			data := map[string]interface{}{
				"motor":        motor,
				"nama_lengkap": session.Values["nama_lengkap"],
			}
			temp := template.Must(template.ParseFiles(filepath.Join("templates", "motor.html")))
			temp.Execute(response, data)
		}
	}
}

func Tambah(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		temp := template.Must(template.ParseFiles(filepath.Join("templates", "tambahmotor.html")))
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {
		request.ParseForm()
		var motor entities.Motor
		motor.Merek = request.Form.Get("merek")
		motor.Tipe = request.Form.Get("tipe")
		motor.JenisMotor = request.Form.Get("jenis_motor")
		motor.TahunProduksi = request.Form.Get("tahun_produksi")
		motor.Warna = request.Form.Get("warna")
		motor.Stok = request.Form.Get("stok")
		motor.Harga = request.Form.Get("harga")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(motor)
		if vErrors != nil {
			data["validation"] = vErrors
		} else {
			if models.NewMotorModel().Create(motor) {
				data["pesan"] = "Data Berhasil Disimpan"
			} else {
				data["pesan"] = "Data Gagal Disimpan"
			}
		}

		temp := template.Must(template.ParseFiles(filepath.Join("templates", "tambahmotor.html")))
		temp.Execute(response, data)

	}
}

func Edit(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var motor entities.Motor
		MotorModel.Find(id, &motor)

		data := map[string]interface{}{
			"motor": motor,
		}

		temp := template.Must(template.ParseFiles(filepath.Join("templates", "editmotor.html")))
		temp.Execute(response, data)
	} else if request.Method == http.MethodPost {
		request.ParseForm()
		var motor entities.Motor
		motor.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		motor.Merek = request.Form.Get("merek")
		motor.Tipe = request.Form.Get("tipe")
		motor.JenisMotor = request.Form.Get("jenis_motor")
		motor.TahunProduksi = request.Form.Get("tahun_produksi")
		motor.Warna = request.Form.Get("warna")
		motor.Stok = request.Form.Get("stok")
		motor.Harga = request.Form.Get("harga")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(motor)
		if vErrors != nil {
			data["motor"] = motor
			data["validation"] = vErrors
		} else {
			if err := models.NewMotorModel().Update(motor); err != nil {
				data["pesan"] = "Data Gagal Diperbarui"
			} else {
				data["pesan"] = "Data Berhasil Diperbarui"
			}
		}

		temp := template.Must(template.ParseFiles(filepath.Join("templates", "editmotor.html")))
		temp.Execute(response, data)

	}
}

func Hapus(response http.ResponseWriter, request *http.Request) {
	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	MotorModel.Delete(id)

	http.Redirect(response, request, "/motor", http.StatusSeeOther)
}

func Sewa(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		customer, _ := customerModel.FindAll()
		motor, _ := MotorModel.FindAll()
		data := map[string]interface{}{
			"customer": customer,
			"motor":    motor,
		}

		temp := template.Must(template.ParseFiles(filepath.Join("templates", "sewa.html")))
		temp.Execute(response, data)
	} else if request.Method == http.MethodPost {
		request.ParseForm()
		var sewa entities.Sewa
		sewa.CustomerId, _ = strconv.ParseInt(request.Form.Get("id_customer"), 10, 64)
		sewa.MotorId, _ = strconv.ParseInt(request.Form.Get("id_motor"), 10, 64)
		sewa.TanggalSewa = request.Form.Get("tanggal_sewa")
		sewa.TanggalKembali = request.Form.Get("tanggal_kembali")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(sewa)
		if vErrors != nil {
			customer, _ := customerModel.FindAll()
			motor, _ := MotorModel.FindAll()
			data["validation"] = vErrors
			data["customer"] = customer
			data["motor"] = motor
		} else {
			if models.NewSewaModel().Create(sewa) {
				MotorModel := models.NewMotorModel()
				if err := MotorModel.UpdateStok(sewa.MotorId); err != nil {
					data["pesan"] = "Motor Berhasil Disewa, Tetapi Terdapat Kesalahan"
				} else {
					data["pesan"] = "Motor Berhasil Disewa"
				}
			} else {
				data["pesan"] = "Motor Gagal Disewa"
			}
		}
		temp := template.Must(template.ParseFiles(filepath.Join("templates", "sewa.html")))
		temp.Execute(response, data)
	}
}
