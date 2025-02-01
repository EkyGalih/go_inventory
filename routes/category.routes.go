package routes

import (
	"inventaris/controllers/categorycontroller"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine) {
	r.GET("/addons/kategori", categorycontroller.Index)
	r.POST("/addons/kategori/store", categorycontroller.Add)
	r.POST("/addons/kategori/update/:id", categorycontroller.Edit)
	r.GET("/addons/kategori/delete/:id", categorycontroller.Delete)
}