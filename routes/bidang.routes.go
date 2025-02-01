package routes

import (
	"inventaris/controllers/bidangcontroller"

	"github.com/gin-gonic/gin"
)

func BidangRoutes(r *gin.Engine) {
	r.GET("/addons/bidang", bidangcontroller.Index)
	r.GET("/addons/bidang/create", bidangcontroller.AddBidang)
	r.POST("/addons/bidang/store", bidangcontroller.AddBidang)
	r.GET("/addons/bidang/edit/:id", bidangcontroller.EditBidang)
	r.POST("/addons/bidang/update/:id", bidangcontroller.EditBidang)
	r.GET("/addons/bidang/delete/:id", bidangcontroller.DeleteBidang)
}