package main

import (
	"fmt"
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
	// config.ConnectDB()

	r := gin.Default()

	// serve static file from public folder
	r.Static("/assets", "./public/assets")
	r.Static("/uploads", "./public/uploads")

	// Inisialisasi session store
	store := cookie.NewStore([]byte("oasioyd8iautsdbiutasdutasduy"))
	r.Use(sessions.Sessions("mysession", store))

	routes.RoutesList(r)
	
	log.Println("Silahkan akses halaman http://localhost:8080")
	r.Run(":8080")
}
