package main

import (
	"Tubes_PBW/controllers/customercontroller"
	"Tubes_PBW/controllers/motorcontroller"
	"Tubes_PBW/controllers/usercontroller"
	"net/http"
)

func main() {

	http.HandleFunc("/", motorcontroller.Dashboard)
	
	http.HandleFunc("/customer", customercontroller.Index)
	http.HandleFunc("/customer/tambah", customercontroller.Tambah)
	http.HandleFunc("/customer/edit", customercontroller.Edit)
	http.HandleFunc("/customer/hapus", customercontroller.Hapus)

	http.HandleFunc("/motor", motorcontroller.Index)
	http.HandleFunc("/motor/tambah", motorcontroller.Tambah)
	http.HandleFunc("/motor/edit", motorcontroller.Edit)
	http.HandleFunc("/motor/hapus", motorcontroller.Hapus)


	http.HandleFunc("/login", usercontroller.Login)
	http.HandleFunc("/logout", usercontroller.Logout)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8787", nil)
}