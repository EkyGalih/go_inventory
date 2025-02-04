package routes

import (
	"fmt"
	"inventaris/controllers/homecontroller"

	"github.com/gin-gonic/gin"
)

func RoutesList(r *gin.Engine) {
	fmt.Println("Routes Loaded")
	// 1. Homepage
	r.GET("/", homecontroller.Welcome)
	
	// 2. categories
	CategoryRoutes(r)

	// 3. Tipe Aset
	TipeRoutes(r)

	// 4. Aset Tetap
	AsetRoutes(r)
	
	// 8. Bidang
	BidangRoutes(r)	
	
	// 9. Pegawai
	PegawaiRoutes(r)
}