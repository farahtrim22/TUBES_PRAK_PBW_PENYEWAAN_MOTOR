package customercontroller

import (
	"Tubes_PBW/entities"
	"Tubes_PBW/libraries"
	"Tubes_PBW/models"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"
)

var validation = libraries.NewValidation()
var CustomerModel = models.NewCustomerModel()

func Index(response http.ResponseWriter, request *http.Request) {
	customer, _ := CustomerModel.FindAll()

	data := map[string]interface{}{
		"customer": customer,
	}

	temp := template.Must(template.ParseFiles(filepath.Join("templates", "customer.html")))
	temp.Execute(response, data)
}

func Tambah(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		temp := template.Must(template.ParseFiles(filepath.Join("templates", "tambahcustomer.html")))
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {
		request.ParseForm()
		var customer entities.Customer
		customer.NamaLengkap = request.Form.Get("nama_lengkap")
		customer.NIK = request.Form.Get("nik")
		customer.JenisKelamin = request.Form.Get("jenis_kelamin")
		customer.Alamat = request.Form.Get("alamat")
		customer.NomorTelepon = request.Form.Get("nomor_telepon")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(customer)
		if vErrors != nil {
			data["validation"] = vErrors
		} else {
			if models.NewCustomerModel().Create(customer) {
				data["pesan"] = "Data Berhasil Disimpan"
			} else {
				data["pesan"] = "Data Gagal Disimpan"
			}
		}

		temp := template.Must(template.ParseFiles(filepath.Join("templates", "tambahcustomer.html")))
		temp.Execute(response, data)

	}
}

func Edit(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var customer entities.Customer
		CustomerModel.Find(id, &customer)

		data := map[string]interface{}{
			"customer": customer,
		}

		temp := template.Must(template.ParseFiles(filepath.Join("templates", "editcustomer.html")))
		temp.Execute(response, data)
	} else if request.Method == http.MethodPost {
		request.ParseForm()
		var customer entities.Customer
		customer.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		customer.NamaLengkap = request.Form.Get("nama_lengkap")
		customer.NIK = request.Form.Get("nik")
		customer.JenisKelamin = request.Form.Get("jenis_kelamin")
		customer.Alamat = request.Form.Get("alamat")
		customer.NomorTelepon = request.Form.Get("nomor_telepon")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(customer)
		if vErrors != nil {
			data["customer"] = customer
			data["validation"] = vErrors
		} else {
			if err := models.NewCustomerModel().Update(customer); err != nil {
				data["pesan"] = "Data Gagal Diperbarui"
			} else {
				data["pesan"] = "Data Berhasil Diperbarui"
			}
		}

		temp := template.Must(template.ParseFiles(filepath.Join("templates", "editcustomer.html")))
		temp.Execute(response, data)

	}
}

func Hapus(response http.ResponseWriter, request *http.Request) {
	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	CustomerModel.Delete(id)

	http.Redirect(response, request, "/", http.StatusSeeOther)
}
