package main

import (
	"inventaris/config"
	"inventaris/controllers/asettikcontroller"
	"inventaris/controllers/categorycontroller"
	"inventaris/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// serve static file from public folder
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	// 1. Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// 2. categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	// 3. aset tik
	http.HandleFunc("/aset_tik", asettikcontroller.Index)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}