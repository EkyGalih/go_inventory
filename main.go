package main

import (
	"inventaris/config"
	"inventaris/controllers/asethabispakaicontroller"
	"inventaris/controllers/asettikcontroller"
	"inventaris/controllers/categorycontroller"
	"inventaris/controllers/homecontroller"
	"inventaris/controllers/lokasiasetcontroller"
	"inventaris/controllers/pemeliharaanaset"
	"inventaris/controllers/tipeasetcontroller"
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
	http.HandleFunc("/addons/kategori", categorycontroller.Index)
	http.HandleFunc("/addons/kategori/add", categorycontroller.Add)
	http.HandleFunc("/addons/kategori/edit", categorycontroller.Edit)
	http.HandleFunc("/addons/kategori/delete", categorycontroller.Delete)

	// 3. aset tik
	http.HandleFunc("/aset/aset-tik", asettikcontroller.Index)
	http.HandleFunc("/aset/aset-tik/add", asettikcontroller.Add)
	http.HandleFunc("/aset/aset-tik/edit", asettikcontroller.Edit)
	http.HandleFunc("/aset/aset-tik/delete", asettikcontroller.Delete)

	http.HandleFunc("/aset/habis-pakai", asethabispakaicontroller.Index)
	http.HandleFunc("/aset/habis-pakai/add", asethabispakaicontroller.Add)
	http.HandleFunc("/aset/habis-pakai/edit", asethabispakaicontroller.Edit)
	http.HandleFunc("/aset/habis-pakai/delete", asethabispakaicontroller.Delete)

	// 4 tipe aset
	http.HandleFunc("/addons/tipe", tipeasetcontroller.Index)
	http.HandleFunc("/addons/tipe/add", tipeasetcontroller.Add)
	http.HandleFunc("/addons/tipe/update", tipeasetcontroller.Update)
	http.HandleFunc("/addons/tipe/delete", tipeasetcontroller.Delete)

	// 5. Pemeliharaan
	http.HandleFunc("/pemeliharaan", pemeliharaanasetcontroller.Index)
	http.HandleFunc("/pemeliharaan/add", pemeliharaanasetcontroller.Add)
	http.HandleFunc("/pemeliharaan/edit", pemeliharaanasetcontroller.Edit)
	http.HandleFunc("/pemeliharaan/status", pemeliharaanasetcontroller.StatusUpdate)
	http.HandleFunc("/pemeliharaan/path", pemeliharaanasetcontroller.GetGambar)

	// 6. Lokasi aset
	http.HandleFunc("/lokasi-aset", lokasiasetcontroller.Index)
	http.HandleFunc("/lokasi-aset/add", lokasiasetcontroller.Add)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}