package asettikcontroller

import (
	"fmt"
	"inventaris/entities"
	"inventaris/helpers"
	"inventaris/models/asettikmodel"
	"inventaris/models/categorymodel"
	"io"
	"net/http"
	"os"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	aset_tiks := asettikmodel.GetAll()
	data := map[string]any{
		"Title":     "Aset TIK",
		"aset_tiks": aset_tiks,
	}

	helpers.RenderTemplate(w, "aset_tik/index.html", data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		categories := categorymodel.GetAll()
		data := map[string]any{
			"Title":      "Add Aset TIK",
			"categories": categories,
		}
		helpers.RenderTemplate(w, "aset_tik/create.html", data)
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
			http.Error(w, "Unable to retrieve file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// save the file
		path := "/public/uploads/aset/"
		fileName := helpers.RandString(30) + ".jpg"
		filePath := path + fileName

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

		aset_tik.Kode_Aset = r.FormValue("kode_aset")
		aset_tik.Nama_Aset = r.FormValue("nama_aset")
		aset_tik.Merek = r.FormValue("merek")
		aset_tik.Model = r.FormValue("model")
		aset_tik.Serial_Number = r.FormValue("serial_number")
		deskripsi := r.FormValue("deskripsi")
		aset_tik.Deskripsi = &deskripsi
		aset_tik.Kategori_id = r.FormValue("kategori_id")
		aset_tik.Tanggal_Perolehan, _ = time.Parse("2006-01-02", r.FormValue("tanggal_perolehan"))
		aset_tik.Status = r.FormValue("status")
		aset_tik.Nilai, _ = helpers.ParseCurrencyToFloat(r.FormValue("nilai"))
		aset_tik.Jumlah = r.FormValue("jumlah")
		keterangan := r.FormValue("keterangan")
		aset_tik.Keterangan = &keterangan
		aset_tik.Path = &filePath
		aset_tik.Gambar = &fileName
		aset_tik.Created_At = time.Now()
		aset_tik.Updated_At = time.Now()

		success, err := asettikmodel.Create(aset_tik)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if success {
			http.Redirect(w, r, "/aset_tik", http.StatusSeeOther)
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
		aset_tik, err := asettikmodel.Detail(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		data := map[string]interface{}{
			"Title":        "Edit Aset TIK",
			"aset_tik":     aset_tik,
			"categories":   categories,
			"SelectedAset": aset_tik.Kategori_id, // untuk selected kategori aset
		}

		helpers.RenderTemplate(w, "aset_tik/edit.html", data)
	}

	if r.Method == http.MethodPost {
		idString := r.FormValue("id")
		if idString == "" {
			http.Error(w, "Missing id parameter", http.StatusBadRequest)
			return
		}

		var aset_tik entities.AsetTik
		oldFilePath, err := asettikmodel.Detail(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// check if a new file is uploaded
		fmt.Println(r.FormValue("gambar"))
		fmt.Println(oldFilePath)

		

		aset_tik.Kode_Aset = r.FormValue("kode_aset")
		aset_tik.Nama_Aset = r.FormValue("nama_aset")
		aset_tik.Merek = r.FormValue("merek")
		aset_tik.Model = r.FormValue("model")
		aset_tik.Serial_Number = r.FormValue("serial_number")
		deskripsi := r.FormValue("deskripsi")
		aset_tik.Deskripsi = &deskripsi
		aset_tik.Kategori_id = r.FormValue("kategori_id")
		aset_tik.Tanggal_Perolehan, _ = time.Parse("2006-01-02", r.FormValue("tanggal_perolehan"))
		aset_tik.Status = r.FormValue("status")
		aset_tik.Nilai, _ = helpers.ParseCurrencyToFloat(r.FormValue("nilai"))
		aset_tik.Jumlah = r.FormValue("jumlah")
		keterangan := r.FormValue("keterangan")
		aset_tik.Keterangan = &keterangan
		aset_tik.Created_At = time.Now()
		aset_tik.Updated_At = time.Now()

		if ok := asettikmodel.Update(idString, aset_tik); !ok {
			http.Error(w, "Failed to update aset_tik", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/aset-tik", http.StatusSeeOther)
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

	http.Redirect(w, r, "/aset-tik", http.StatusSeeOther)
}
