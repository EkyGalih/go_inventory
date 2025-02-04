package routes

import (
	"inventaris/controllers/pegawaicontroller"

	"github.com/gin-gonic/gin"
)

func PegawaiRoutes(r *gin.Engine) {
	r.GET("/addons/pegawai", pegawaicontroller.Index)
	r.GET("/addons/pegawai/create", pegawaicontroller.Add)
	r.POST("/addons/pegawai/store", pegawaicontroller.Store)
	r.GET("/addons/pegawai/edit/:id", pegawaicontroller.Edit)
	r.POST("/addons/pegawai/update/:id", pegawaicontroller.Update)
	r.GET("/addons/pegawai/delete/:id", pegawaicontroller.Delete)
}
