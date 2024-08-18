package homecontroller

import (
	"inventaris/helpers/helpers"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	path := map[string]string{
		"menu": "dashboard",
	}
	data := map[string]any{
		"Title": "Dashboard",
		"path":  path,
	}
	helpers.RenderTemplate(w, "home/index.html", data)
}
