package main

import (
	"fmt"
	"inventaris/config"
	"inventaris/routes"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("\033[32m====================================================")
	fmt.Println("\033[32m  SELAMAT DATANG DI APLIKASI INVENTARIS ASSET  ")
	fmt.Println("\033[32m====================================================")
	fmt.Println("\033[0m")
	config.ConnectDB()

	r := gin.Default()

	// serve static file from public folder
	r.Static("/assets", "./public/assets")
	r.Static("/uploads", "./public/uploads")

	// Inisialisasi session store
	store := cookie.NewStore([]byte("oasioyd8iautsdbiutasdutasduy"))
	r.Use(sessions.Sessions("mysession", store))

	// 2. categories
	// r.GET("/addons/kategori", categorycontroller.Index)
	// r.GET("/addons/kategori/add", categorycontroller.Add)
	// r.GET("/addons/kategori/edit", categorycontroller.Edit)
	// r.GET("/addons/kategori/delete", categorycontroller.Delete)

	// // 3. aset tik
	// r.GET("/aset/aset-tik", asettikcontroller.Index)
	// r.GET("/aset/aset-tik/add", asettikcontroller.Add)
	// r.GET("/aset/aset-tik/edit", asettikcontroller.Edit)
	// r.GET("/aset/aset-tik/distribusi", asettikcontroller.Distribusi)
	// r.GET("/aset/aset-tik/delete", asettikcontroller.Delete)

	// r.GET("/aset/habis-pakai", asethabispakaicontroller.Index)
	// r.GET("/aset/habis-pakai/add", asethabispakaicontroller.Add)
	// r.GET("/aset/habis-pakai/edit", asethabispakaicontroller.Edit)
	// r.GET("/aset/habis-pakai/delete", asethabispakaicontroller.Delete)

	// // 4 tipe aset
	// r.GET("/addons/tipe", tipeasetcontroller.Index)
	// r.GET("/addons/tipe/add", tipeasetcontroller.Add)
	// r.GET("/addons/tipe/update", tipeasetcontroller.Update)
	// r.GET("/addons/tipe/delete", tipeasetcontroller.Delete)

	// // 5. Pemeliharaan
	// r.GET("/pemeliharaan", pemeliharaanasetcontroller.Index)
	// r.GET("/pemeliharaan/add", pemeliharaanasetcontroller.Add)
	// r.GET("/pemeliharaan/edit", pemeliharaanasetcontroller.Edit)
	// r.GET("/pemeliharaan/status", pemeliharaanasetcontroller.StatusUpdate)
	// r.GET("/pemeliharaan/path", pemeliharaanasetcontroller.GetGambar)

	// // 6. Lokasi aset
	// r.GET("/lokasi-aset", lokasiasetcontroller.Index)
	// r.GET("/lokasi-aset/add", lokasiasetcontroller.Add)
	// r.GET("/lokasi-aset/edit", lokasiasetcontroller.Edit)
	// // r.GET("/lokasi-aset/daftar", lokasiasetcontroller.AsetPegawai)

	// // 7. Riwayat Aset
	// r.GET("/riwayat-aset", riwayatasetcontroller.Index)
	// r.GET("/riwayat-aset/logs", riwayatasetcontroller.Show)

	// // 8. Bidang
	// r.GET("/addons/bidang", bidangcontroller.Index)
	// r.GET("/addons/bidang/create", bidangcontroller.AddBidang)
	// r.GET("/addons/bidang/edit", bidangcontroller.EditBidang)

	routes.RoutesList(r)

	log.Println("Silahkan akses halaman http://localhost:8080")
	r.Run(":8080")
}
