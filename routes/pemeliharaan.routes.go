package routes

import (
	"inventaris/controllers/pemeliharaancontroller"

	"github.com/gin-gonic/gin"
)

func PemeliharaanRoutes(r *gin.Engine) {
	r.GET("/pemeliharaan", pemeliharaancontroller.Index)
	r.GET("/pemeliharaan/create", pemeliharaancontroller.Add)
	r.POST("/pemeliharaan/store", pemeliharaancontroller.Store)
	r.GET("/pemeliharaan/edit/:id", pemeliharaancontroller.Edit)
	r.POST("/pemeliharaan/update/:id", pemeliharaancontroller.Update)
	r.GET("/pemeliharaan/delete/:id", pemeliharaancontroller.Delete)
	r.GET("/pemeliharaan/path/:id", pemeliharaancontroller.GetGambar)
}