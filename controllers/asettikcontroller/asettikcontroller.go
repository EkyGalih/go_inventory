package asettikcontroller

import (
	"inventaris/entities"
	"inventaris/helpers"
	"inventaris/models/asettikmodel"
	"inventaris/models/categorymodel"
	"inventaris/models/tipemodel"
	"io"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 5
	}

	aset_tiks, err := asettikmodel.GetAll(page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	totalRows, err := asettikmodel.GetTotalRows()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	totalPages := int(math.Ceil(float64(totalRows) / float64(limit)))

	distribusi := helpers.GetDistribusi(aset_tiks)
	if distribusi == nil {
		distribusi = make(map[string]int)
	}

	// untuk show menu dan active sub menu
	path := map[string]string{
		"menu":    "aset",
		"subMenu": "aset-tik",
	}

	data := map[string]any{
		"Title":      "Aset TIK",
		"path":       path,
		"Page":       page,
		"TotalPages":  totalPages,
		"TotalRows":  totalRows,
		"Limit":      limit,
		"aset_tiks":  aset_tiks,
		"distribusi": distribusi,
	}

	helpers.RenderTemplate(w, "/aset/aset_tik/index.html", data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		categories := categorymodel.GetAll()
		tipes := tipemodel.GetAll()

		// untuk show menu dan active sub menu
		path := map[string]string{
			"menu":    "aset",
			"subMenu": "aset-tik",
		}

		data := map[string]any{
			"Title":      "Add Aset TIK",
			"path":       path,
			"categories": categories,
			"tipes":      tipes,
		}
		helpers.RenderTemplate(w, "/aset/aset_tik/create.html", data)
	}

	if r.Method == http.MethodPost {
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var aset_tik entities.AsetTik

		// handle file upload
		file, _, err := r.FormFile("gambar")
		if err != nil {
			http.Error(w, "Unable to retrieve file "+err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// save the file
		path := "./public/uploads/aset/tetap/"
		fileName := helpers.RandString(30) + ".jpg"
		filePath := filepath.Join(path, fileName)
		dbPath := strings.ReplaceAll(filepath.Join("/public/uploads/aset/tetap/", fileName), "\\", "/")

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

		aset_tik.Jenis_Aset = r.FormValue("jenis_aset")
		aset_tik.Kode_Aset = r.FormValue("kode_aset")
		aset_tik.Nama_Aset = r.FormValue("nama_aset")
		aset_tik.Merek = r.FormValue("merek")
		aset_tik.Model = r.FormValue("model")
		aset_tik.Serial_Number = r.FormValue("serial_number")
		deskripsi := r.FormValue("deskripsi")
		aset_tik.Deskripsi = &deskripsi
		aset_tik.Kategori_id = r.FormValue("kategori_id")
		aset_tik.Tipe_id = r.FormValue("tipe_id")
		aset_tik.Tanggal_Perolehan, _ = time.Parse("2006-01-02", r.FormValue("tanggal_perolehan"))
		aset_tik.Status = r.FormValue("status")
		aset_tik.Nilai, _ = helpers.ParseCurrencyToFloat(r.FormValue("nilai"))
		aset_tik.Jumlah, err = strconv.ParseFloat(r.FormValue("jumlah"), 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		keterangan := r.FormValue("keterangan")
		aset_tik.Keterangan = &keterangan
		aset_tik.Path = &dbPath
		aset_tik.Gambar = &fileName
		aset_tik.Created_At = time.Now()
		aset_tik.Updated_At = time.Now()

		success, err := asettikmodel.Create(aset_tik)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if success {
			http.Redirect(w, r, "/aset/aset-tik", http.StatusSeeOther)
		} else {
			http.Error(w, "Failed to create aset_tik", http.StatusInternalServerError)
		}
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idString := r.URL.Query().Get("id")
		if idString == "" {
			http.Error(w, "Missing id parameter", http.StatusBadRequest)
			return
		}

		categories := categorymodel.GetAll()
		tipes := tipemodel.GetAll()
		aset_tik, err := asettikmodel.Detail(idString)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		// untuk show menu dan active sub menu
		path := map[string]string{
			"menu":    "aset",
			"subMenu": "aset-tik",
		}

		data := map[string]interface{}{
			"Title":        "Edit Aset TIK",
			"path":         path,
			"aset_tik":     aset_tik,
			"categories":   categories,
			"tipes":        tipes,
			"SelectedTipe": aset_tik.Tipe_id,
			"SelectedAset": aset_tik.Kategori_id, // untuk selected kategori aset
		}

		helpers.RenderTemplate(w, "aset/aset_tik/edit.html", data)
	}

	if r.Method == http.MethodPost {
		idString := r.FormValue("id")
		if idString == "" {
			http.Error(w, "Missing id parameter", http.StatusBadRequest)
			return
		}

		aset, err := asettikmodel.Detail(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		oldFilePath := filepath.Join(".", *aset.Path)
		var aset_tik entities.AsetTik
		var dbPath, fileName string

		// check if a new file is uploaded
		file, _, err := r.FormFile("gambar")
		if err != nil {
			if err == http.ErrMissingFile {
				// Tidak ada file baru, lanjutkan tanpa mengubah file lama
				file = nil
				dbPath = *aset.Path
				fileName = *aset.Gambar
			} else {
				http.Error(w, "Error retrieving file", http.StatusInternalServerError)
				return
			}
		} else {
			// Hapus file lama jika ada
			if _, err := os.Stat(oldFilePath); err == nil {
				err := os.Remove(oldFilePath)
				if err != nil {
					http.Error(w, "Failed to delete old file", http.StatusInternalServerError)
					return
				}
			}

			// Simpan file baru
			path := "./public/uploads/aset/tetap/"
			fileName = helpers.RandString(30) + ".jpg"
			newFilePath := filepath.Join(path, fileName)
			dbPath = strings.ReplaceAll(filepath.Join("/public/uploads/aset/tetap/", fileName), "\\", "/")

			out, err := os.Create(newFilePath)

			if err != nil {
				http.Error(w, "Failed to create new file", http.StatusInternalServerError)
				return
			}
			defer out.Close()

			_, err = io.Copy(out, file)
			if err != nil {
				http.Error(w, "Failed to save new file", http.StatusInternalServerError)
				return
			}
			defer file.Close()
		}

		aset_tik.Nama_Aset = r.FormValue("nama_aset")
		aset_tik.Jenis_Aset = r.FormValue("jenis_aset")
		aset_tik.Merek = r.FormValue("merek")
		aset_tik.Model = r.FormValue("model")
		aset_tik.Tanggal_Perolehan, _ = time.Parse("2006-01-02", r.FormValue("tanggal_perolehan"))
		aset_tik.Nilai, _ = helpers.ParseCurrencyToFloat(r.FormValue("nilai"))
		deskripsi := r.FormValue("deskripsi")
		aset_tik.Deskripsi = &deskripsi
		aset_tik.Path = &dbPath
		aset_tik.Gambar = &fileName
		aset_tik.Kode_Aset = r.FormValue("kode_aset")
		aset_tik.Kategori_id = r.FormValue("kategori_id")
		aset_tik.Tipe_id = r.FormValue("tipe_id")
		aset_tik.Serial_Number = r.FormValue("serial_number")
		aset_tik.Status = r.FormValue("status")
		aset_tik.Jumlah, err = strconv.ParseFloat(r.FormValue("jumlah"), 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		keterangan := r.FormValue("keterangan")
		aset_tik.Keterangan = &keterangan
		aset_tik.Created_At = time.Now()
		aset_tik.Updated_At = time.Now()

		success, err := asettikmodel.Update(idString, aset_tik)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if success {
			http.Redirect(w, r, "/aset/aset-tik", http.StatusSeeOther)
		} else {
			http.Error(w, "Failed to update aset_tik", http.StatusInternalServerError)
		}
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	if idString == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	if err := asettikmodel.Delete(idString); err != nil {
		http.Error(w, "Failed to delete aset_tik", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/aset/aset-tik", http.StatusSeeOther)
}
