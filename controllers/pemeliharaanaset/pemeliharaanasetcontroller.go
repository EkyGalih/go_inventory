package pemeliharaanasetcontroller

import (
	"inventaris/helpers"
	"inventaris/models/PemeliharaanModel"
	"inventaris/models/asettikmodel"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	pemeliharaan := pemeliharaanmodel.GetAll()
	path := map[string]string{
		"menu": "pemeliharaan",
	}
	data := map[string]any{
		"Title":        "Pemeliharaan Aset",
		"path":         path,
		"pemeliharaan": pemeliharaan,
	}
	helpers.RenderTemplate(w, "pemeliharaan/index.html", data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		pemeliharaan := pemeliharaanmodel.GetAll()
		aset_tik := asettikmodel.GetAll()

		path := map[string]string{
			"menu": "pemeliharaan",
		}
		data := map[string]any{
			"Title":        "Pemeliharaan Aset",
			"path":         path,
			"pemeliharaan": pemeliharaan,
			"aset_tik":     aset_tik,
		}
		helpers.RenderTemplate(w, "pemeliharaan/create.html", data)
	}
}

func GetGambar(w http.ResponseWriter, r *http.Request) {
    // Mendapatkan ID dari query string
    idString := r.URL.Query().Get("id")
    if idString == "" {
        http.Error(w, "Parameter id tidak ditemukan", http.StatusBadRequest)
        return
    }

    // Mendapatkan detail aset berdasarkan ID
    aset, err := asettikmodel.Detail(idString)
    if err != nil {
        http.Error(w, "Gagal mendapatkan detail aset: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Pastikan path gambar ditemukan
    if aset.Path == nil || *aset.Path == "" {
        http.Error(w, "Path gambar tidak ditemukan", http.StatusNotFound)
        return
    }

    // Mengatur header Content-Type sebagai text/plain untuk mengirim path sebagai string
    w.Header().Set("Content-Type", "text/plain")

    // Menulis path gambar ke dalam respon
    _, err = w.Write([]byte(*aset.Path))
    if err != nil {
        http.Error(w, "Gagal menulis respons: "+err.Error(), http.StatusInternalServerError)
    }
}

