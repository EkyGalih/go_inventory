package homecontroller

import (
	"inventaris/helpers"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Title": "Dashboard",
	}
	helpers.RenderTemplate(w, "home/index.html", data)
}
