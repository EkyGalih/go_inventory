package pemeliharaanasetcontroller

import (
	"inventaris/entities"
	"inventaris/helpers/helpers"
	"inventaris/models/PemeliharaanModel"
	"inventaris/models/asettikmodel"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
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

	if r.Method == http.MethodPost {
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var pemeliharaan entities.Pemeliharaan
		var dbPath string

		file, _, err := r.FormFile("nota")

		if err != nil {
			if err == http.ErrMissingFile {
				dbPath = ""
			} else {
				http.Error(w, "Error Retrieving the File: "+err.Error(), http.StatusInternalServerError)
			}
		} else {
			defer file.Close()

			//save nota
			path := "./public/uploads/pemeliharaan/nota"
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				http.Error(w, "failed to create directory", http.StatusInternalServerError)
				return
			}
			filename := helpers.RandString(20) + ".pdf"
			filePath := filepath.Join(path, filename)
			dbPath = strings.ReplaceAll(filepath.Join("/public/uploads/pemeliharaan/nota", filename), "\\", "/")

			dest, err := os.Create(filePath)
			if err != nil {
				http.Error(w, "Unable to save file: "+err.Error(), http.StatusInternalServerError)
				return
			}
			defer dest.Close()

			_, err = io.Copy(dest, file)
			if err != nil {
				http.Error(w, "Unable to save file: "+err.Error(), http.StatusInternalServerError)
				return
			}
		}

		pemeliharaan.Aset_id = r.FormValue("aset_id")
		pemeliharaan.Tanggal_Pemeliharaan, _ = time.Parse("2006-01-02", r.FormValue("tanggal_pemeliharaan"))
		kerusakan := r.FormValue("kerusakan")
		perbaikan := r.FormValue("perbaikan")
		keterangan := r.FormValue("keterangan")
		pemeliharaan.Kerusakan = &kerusakan
		pemeliharaan.Perbaikan = &perbaikan
		pemeliharaan.Keterangan = &keterangan
		pemeliharaan.Status = "Proses"
		pemeliharaan.Nota = &dbPath
		pemeliharaan.Biaya, err = helpers.ParseCurrencyToFloat(r.FormValue("biaya"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		pemeliharaan.Created_At = time.Now()
		pemeliharaan.Updated_At = time.Now()

		success, err := pemeliharaanmodel.Create(pemeliharaan)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if success {
			http.Redirect(w, r, "/pemeliharaan", http.StatusSeeOther)
		} else {
			http.Error(w, "Failed to create pemeliharaan", http.StatusInternalServerError)
		}
	}

}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idString := r.URL.Query().Get("id")
		if idString == "" {
			http.Error(w, "Parameter id tidak ditemukan", http.StatusBadRequest)
			return
		}

		pemeliharaan, err := pemeliharaanmodel.Detail(idString)
		aset_tik := asettikmodel.GetAll()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		path := map[string]string{
			"menu": "pemeliharaan",
		}

		data := map[string]interface{}{
			"Title":        "Edit Pemeliharaan Aset",
			"path":         path,
			"aset_tik":     aset_tik,
			"SelectedAset": pemeliharaan.Aset_id,
			"pemeliharaan": pemeliharaan,
		}

		helpers.RenderTemplate(w, "pemeliharaan/edit.html", data)
	}

	if r.Method == http.MethodPost {
		idString := r.FormValue("id")
		if idString == "" {
			http.Error(w, "Parameter id tidak ditemukan", http.StatusBadRequest)
			return
		}

		pemeliharaan, err := pemeliharaanmodel.Detail(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		oldFilepath := filepath.Join(".", *pemeliharaan.Nota)
		var data entities.Pemeliharaan
		var dbPath, fileName string

		file, _, err := r.FormFile("nota")
		if err != nil {
			if err == http.ErrMissingFile {
				file = nil
				dbPath = *pemeliharaan.Nota
				fileName = ""
			} else {
				http.Error(w, "erro retrieving file", http.StatusInternalServerError)
				return
			}
		} else {
			// Check if the old file path is not null or empty
			if oldFilepath != "" && *pemeliharaan.Nota != "" {
				// Attempt to delete the old file if it exists
				if _, err := os.Stat(oldFilepath); err == nil {
					err = os.Remove(oldFilepath)
					if err != nil {
						http.Error(w, "failed to delete old file", http.StatusInternalServerError)
						return
					}
				}
			}

			// simpan file baru
			path := "./public/uploads/pemeliharaan/nota/"
			fileName = helpers.RandString(15) + ".pdf"
			newFilePath := filepath.Join(path, fileName)
			dbPath = strings.ReplaceAll(filepath.Join("/public/uploads/pemeliharaan/nota/", fileName), "\\", "/")

			out, err := os.Create(newFilePath)

			if err != nil {
				http.Error(w, "failed to create new file", http.StatusInternalServerError)
				return
			}
			defer out.Close()

			_, err = io.Copy(out, file)
			if err != nil {
				http.Error(w, "failed to save new file", http.StatusInternalServerError)
				return
			}
			defer file.Close()

		}

		data.Aset_id = r.FormValue("aset_id")
		data.Tanggal_Pemeliharaan, _ = time.Parse("2006-01-02", r.FormValue("tanggal_pemeliharaan"))
		kerusakan := r.FormValue("kerusakan")
		perbaikan := r.FormValue("perbaikan")
		keterangan := r.FormValue("keterangan")
		data.Kerusakan = &kerusakan
		data.Perbaikan = &perbaikan
		data.Keterangan = &keterangan
		data.Status = "Proses"
		data.Nota = &dbPath
		data.Biaya, err = helpers.ParseCurrencyToFloat(r.FormValue("biaya"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data.Created_At = time.Now()
		data.Updated_At = time.Now()

		success, err := pemeliharaanmodel.Update(idString, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if success {
			http.Redirect(w, r, "/pemeliharaan", http.StatusSeeOther)
		} else {
			http.Error(w, "Failed to update pemeliharaan", http.StatusInternalServerError)
		}
	}
}

func StatusUpdate(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan ID dari query string
	idString := r.URL.Query().Get("id")
	if idString == "" {
		http.Error(w, "Parameter id tidak ditemukan", http.StatusBadRequest)
		return
	}

	// Mendapatkan detail pemeliharaan berdasarkan ID
	pemeliharaan, err := pemeliharaanmodel.Detail(idString)
	if err != nil {
		http.Error(w, "Gagal mendapatkan detail pemeliharaan: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Update status pemeliharaan
	if pemeliharaan.Status == "Selesai" {
		pemeliharaan.Status = "Proses"
	} else {
		pemeliharaan.Status = "Selesai"
		pemeliharaan.Updated_At = time.Now()
	}

	success, err := pemeliharaanmodel.Update(idString, pemeliharaan)
	if err != nil {
		http.Error(w, "Gagal memperbarui status pemeliharaan: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if success {
		http.Redirect(w, r, "/pemeliharaan", http.StatusSeeOther)
	} else {
		http.Error(w, "Gagal memperbarui status pemeliharaan", http.StatusInternalServerError)
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
