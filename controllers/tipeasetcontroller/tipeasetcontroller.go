package tipeasetcontroller

import (
	"fmt"
	"inventaris/entities"
	"inventaris/helpers"
	"inventaris/models/tipemodel"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tipes := tipemodel.GetAll()
	path := map[string]string{
		"menu":    "addons",
		"subMenu": "type",
	}
	data := map[string]any{
		"Title": "Tipe Aset",
		"path":  path,
		"tipes": tipes,
	}
	helpers.RenderTemplate(w, "addons/type/index.html", data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var tipe entities.Tipe

		tipe.Nama_Tipe = r.FormValue("nama_tipe")
		keterangan := r.FormValue("keterangan")
		tipe.Keterangan = &keterangan
		tipe.Created_At = time.Now()
		tipe.Updated_At = time.Now()

		success, err := tipemodel.Create(tipe)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if success {
			http.Redirect(w, r, "/addons/tipe", http.StatusSeeOther)
		} else {
			http.Error(w, "Gagal Membuat tipe aset", http.StatusInternalServerError)
		}
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		idString := r.FormValue("id")
		if idString == "" {
			http.Error(w, "Parameter Id tidak ditemukan", http.StatusBadRequest)
			return
		}

		var tipe entities.Tipe
		tipe.Nama_Tipe = r.FormValue("nama_tipe")
		keterangan := r.FormValue("keterangan")
		tipe.Keterangan = &keterangan
		tipe.Updated_At = time.Now()

		if ok := tipemodel.Update(idString, tipe); !ok {
			http.Error(w, "Gagal Mengupdate tipe aset", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/addons/tipe", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	fmt.Println("idString", idString)
	if idString == "" {
		http.Error(w, "Parameter Id tidak ditemukan", http.StatusBadRequest)
		return
	}

	if err := tipemodel.Delete(idString); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/addons/tipe", http.StatusSeeOther)
}