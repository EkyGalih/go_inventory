package routes

import (
	"inventaris/controllers/tipeasetcontroller"

	"github.com/gin-gonic/gin"
)

func TipeRoutes(r *gin.Engine) {
	r.GET("/addons/tipe", tipeasetcontroller.Index)
	r.POST("/addons/tipe/store", tipeasetcontroller.Add)
	r.POST("/addons/tipe/update/:id", tipeasetcontroller.Edit)
	r.GET("/addons/tipe/delete/:id", tipeasetcontroller.Delete)
}