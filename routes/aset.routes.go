package routes

import (
	"inventaris/controllers/asetcontroller"

	"github.com/gin-gonic/gin"
)

func AsetRoutes(r *gin.Engine) {
	r.GET("/aset", asetcontroller.Index)
	r.GET("/aset/create", asetcontroller.Add)
	r.POST("/aset/store", asetcontroller.Store)
	r.GET("/aset/edit/:id", asetcontroller.Edit)
	r.POST("/aset/update/:id", asetcontroller.Update)
	r.GET("/aset/delete/:id", asetcontroller.Delete)
}
