package asettikcontroller

import (
	"inventaris/helpers"
	"inventaris/models/asettikmodel"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	aset_tiks := asettikmodel.GetAll()
	data := map[string]any{
		"Title": "Aset TIK",
		"aset_tiks": aset_tiks,
	}

	helpers.RenderTemplate(w, "aset_tik/index.html", data)
}