package authcontroller

import (
	"inventaris/helpers/helpers"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	path := map[string]string{
		"menu": "login",
	}
	data := map[string]any{
		"Title": "Login",
		"path":  path,
	}
	helpers.RenderTemplate(w, "auth/login.html", data)
}