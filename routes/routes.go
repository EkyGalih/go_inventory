package routes

import (
	"inventaris/controllers/homecontroller"

	"github.com/gin-gonic/gin"
)

func RoutesList(r *gin.Engine) {
	// 1. Homepage
	r.GET("/", homecontroller.Welcome)
	
	// 2. categories
	CategoryRoutes(r)
	
	// 8. Bidang
	BidangRoutes(r)
}