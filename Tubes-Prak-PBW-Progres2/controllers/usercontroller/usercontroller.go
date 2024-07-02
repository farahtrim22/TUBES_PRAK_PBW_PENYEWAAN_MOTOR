package usercontroller

import (
	"Tubes_PBW/config"
	"Tubes_PBW/entities"
	"Tubes_PBW/models"
	"errors"
	"net/http"
	"path/filepath"
	"text/template"
	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Username 	string
	Password 	string
}

var userModel = models.NewUserModel()

func Login(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		temp := template.Must(template.ParseFiles(filepath.Join("templates", "login.html")))
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {
		//proses login
		request.ParseForm()
		UserInput := &UserInput {
			Username: request.Form.Get("username"),
			Password: request.Form.Get("password"),
		}

		var user entities.User
		userModel.Where(&user, "username", UserInput.Username)

		var massage error
		if user.Username == " " {
			// tidak ada di data base
			massage = errors.New("Username atau Password salah!!!")
		} else {
			// pengecekan password
			errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserInput.Password))
			if errPassword != nil {
				massage = errors.New("Username atau Password salah!!!")
			}
		}

		if massage != nil {

			data := map[string]interface{} {
				"error" : massage,
			}

			temp := template.Must(template.ParseFiles(filepath.Join("templates", "login.html")))
			temp.Execute(response, data)
		} else {
			// ngeset session 

			session, _ := config.Store.Get(request, config.SESSION_ID)

			session.Values["loggedIn"] = true
			session.Values["email"] = user.Email
			session.Values["username"] = user.Username
			session.Values["nama_lengkap"] = user.NamaLengkap

			session.Save(request, response)

			http.Redirect(response, request, "/", http.StatusSeeOther)
		}

	}
}

func Logout(response http.ResponseWriter, request *http.Request) {

	session, _ := config.Store.Get(request, config.SESSION_ID)

	// hapus session
	session.Options.MaxAge = -1
	session.Save(request, response)

	http.Redirect(response, request, "/login", http.StatusSeeOther)


}