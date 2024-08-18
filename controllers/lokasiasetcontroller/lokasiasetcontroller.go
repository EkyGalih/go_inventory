package lokasiasetcontroller

import (
	"fmt"
	"inventaris/entities"
	"inventaris/helpers/helpers"
	"inventaris/helpers/queryhelpers"
	"inventaris/models/asettikmodel"
	"inventaris/models/bidangmodel"
	"inventaris/models/lokasiasetmodel"
	"inventaris/models/pegawaimodel"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	lokasiaset := lokasiasetmodel.GetAll()
	aset := asettikmodel.GetAll()
	bidang := bidangmodel.GetAll()
	pegawai := pegawaimodel.GetALl()

	path := map[string]string{
		"menu": "lokasi-aset",
	}
	aset_pegawai := queryhelpers.CountAsetPegawai(lokasiaset)
	if aset_pegawai == nil {
		aset_pegawai = make(map[string]int)
	}

	data := map[string]any{
		"Title":        "Lokasi Aset",
		"path":         path,
		"lokasiaset":   lokasiaset,
		"aset":         aset,
		"bidang":       bidang,
		"pegawai":      pegawai,
		"aset_pegawai": aset_pegawai,
	}

	helpers.RenderTemplate(w, "lokasi_aset/index.html", data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		aset := asettikmodel.GetAll()
		bidang := bidangmodel.GetAll()
		pegawai := pegawaimodel.GetALl()

		path := map[string]string{
			"menu": "lokasi-aset",
		}
		data := map[string]any{
			"Title":   "Lokasi Aset",
			"path":    path,
			"aset":    aset,
			"bidang":  bidang,
			"pegawai": pegawai,
		}
		helpers.RenderTemplate(w, "lokasi_aset/create.html", data)
	}

	if r.Method == http.MethodPost {
		var lokasiaset entities.LokasiAset

		lokasiaset.Aset_id = r.FormValue("aset_id")
		lokasiaset.Bidang_id = r.FormValue("bidang_id")
		lokasiaset.Pegawai_id = r.FormValue("pegawai_id")
		lokasiaset.Tanggal_Perolehan, _ = time.Parse("2006-01-02", r.FormValue("tanggal_perolehan"))
		tglSelesaiStr := r.FormValue("tanggal_selesai")
		if tglSelesaiStr != "" {
			tgl_selesai, err := time.Parse("2006-01-02", tglSelesaiStr)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			lokasiaset.Tanggal_Selesai = &tgl_selesai
		} else {
			lokasiaset.Tanggal_Selesai = nil
		}
		jenis_pemanfaatan := r.FormValue("jenis_pemanfaatan")
		lokasiaset.Jenis_Pemanfaatan = &jenis_pemanfaatan
		keterangan := r.FormValue("keterangan")
		lokasiaset.Keterangan = &keterangan
		lokasiaset.Created_At = time.Now()
		lokasiaset.Updated_At = time.Now()

		success, err := lokasiasetmodel.Create(lokasiaset)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if success {
			http.Redirect(w, r, "/lokasi-aset", http.StatusSeeOther)
		} else {
			http.Error(w, "Failed to create lokasi aset", http.StatusInternalServerError)
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

		aset := asettikmodel.GetAll()
		bidang := bidangmodel.GetAll()
		pegawai := pegawaimodel.GetALl()
		lokasiaset, err := lokasiasetmodel.Detail(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		path := map[string]string{
			"menu": "lokasi-aset",
		}

		data := map[string]interface{}{
			"Title":             "Edit Lokasi Aset",
			"path":              path,
			"aset":              aset,
			"bidang":            bidang,
			"pegawai":           pegawai,
			"lokasiaset":        lokasiaset,
			"SelectedAset":      lokasiaset.Aset_id,
			"SelectedBidang":    lokasiaset.Bidang_id,
			"SelectedPegawai":   lokasiaset.Pegawai_id,
			"SelectedJenisAset": lokasiaset.Jenis_Pemanfaatan,
		}

		helpers.RenderTemplate(w, "lokasi_aset/edit.html", data)
	}

	if r.Method == http.MethodPost {
		idString := r.FormValue("id")
		if idString == "" {
			http.Error(w, "Parameter id tidak ditemukan", http.StatusBadRequest)
			return
		}
		fmt.Println(r.FormValue("jenis_pemanfaatan"))
		var lokasiaset entities.LokasiAset
		lokasiaset.Aset_id = r.FormValue("aset_id")
		lokasiaset.Bidang_id = r.FormValue("bidang_id")
		lokasiaset.Pegawai_id = r.FormValue("pegawai_id")
		lokasiaset.Tanggal_Perolehan, _ = time.Parse("2006-01-02", r.FormValue("tanggal_perolehan"))
		tglSelesaiStr := r.FormValue("tanggal_selesai")
		if tglSelesaiStr != "" {
			tgl_selesai, err := time.Parse("2006-01-02", tglSelesaiStr)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			lokasiaset.Tanggal_Selesai = &tgl_selesai
		} else {
			lokasiaset.Tanggal_Selesai = nil
		}
		jenis_pemanfaatan := r.FormValue("jenis_pemanfaatan")
		lokasiaset.Jenis_Pemanfaatan = &jenis_pemanfaatan
		keterangan := r.FormValue("keterangan")
		lokasiaset.Keterangan = &keterangan
		lokasiaset.Updated_At = time.Now()

		success, err := lokasiasetmodel.Update(idString, lokasiaset)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if success {
			http.Redirect(w, r, "/lokasi-aset", http.StatusSeeOther)
		} else {
			http.Error(w, "Failed to update lokasi aset", http.StatusInternalServerError)
		}
	}
}

func AsetPegawai(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("pegawai_id")
	if idString == "" {
		http.Error(w, "Parameter id tidak ditemukan", http.StatusBadRequest)
		return
	}

	asetPegawai := lokasiasetmodel.DaftarAset(idString)

	asetIDs := make([]string, len(asetPegawai))
	for i, aset := range asetPegawai {
		asetIDs[i] = aset.Id
	}
	var lokasiaset entities.LokasiAset
	if len(asetIDs) > 0 {
		var err error
		lokasiaset, err = lokasiasetmodel.Detail(asetIDs[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	aset_pegawai := queryhelpers.CountAsetPegawai(asetPegawai)
	if aset_pegawai == nil {
		aset_pegawai = make(map[string]int)
	}

	path := map[string]string{
		"menu": "lokasi-aset",
	}

	data := map[string]interface{}{
		"Title":            "Daftar Aset Pegawai",
		"path":             path,
		"asetPegawai":      asetPegawai,
		"lokasiaset":       lokasiaset,
		"countAsetPegawai": aset_pegawai,
	}

	helpers.RenderTemplate(w, "lokasi_aset/daftar_aset.html", data)

}
