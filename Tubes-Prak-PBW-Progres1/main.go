package main

import (
	"Tubes_PBW/controllers/customercontroller"
	"net/http"
)

func main() {

	http.HandleFunc("/", customercontroller.Index)
	http.HandleFunc("/customer/tambah", customercontroller.Tambah)
	http.HandleFunc("/customer/edit", customercontroller.Edit)
	http.HandleFunc("/customer/hapus", customercontroller.Hapus)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8787", nil)
}
