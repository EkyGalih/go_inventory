package homecontroller

import (
	"inventaris/helpers/helpers"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	path := map[string]string{
		"menu": "dashboard",
	}
	
	data := map[string]any{
		"Title": "Dashboard",
		"path":  path,
	}
	helpers.RenderTemplate(c, "/home/index.html", data)
}
