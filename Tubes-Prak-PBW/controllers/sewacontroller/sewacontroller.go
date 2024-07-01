package sewacontroller

import (
	"Tubes_PBW/config"
	"Tubes_PBW/models"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

var SewaModel = models.NewSewaModel()
var MotorModel = models.NewMotorModel()
var customerModel = models.NewCustomerModel()

func Dashboard(response http.ResponseWriter, request *http.Request) {

	session, _ := config.Store.Get(request, config.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(response, request, "/login", http.StatusSeeOther)
	} else {
		if session.Values["loggedIn"] != true {
			http.Redirect(response, request, "/login", http.StatusSeeOther)
		} else {
			customer, _ := customerModel.FindAll()
			motor, _ := MotorModel.FindAll()
			sewa, _ := SewaModel.FindAll()
			data := map[string]interface{}{
				"sewa":        sewa,
				"customer":    customer,
				"motor":       motor,
				"nama_lengkap": session.Values["nama_lengkap"],
			}
			temp := template.Must(template.ParseFiles(filepath.Join("templates", "index.html")))
			temp.Execute(response, data)
		}
	}
}

func Hapus(response http.ResponseWriter, request *http.Request) {
	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	SewaModel.Delete(id)

	http.Redirect(response, request, "/", http.StatusSeeOther)
}
