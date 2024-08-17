package lokasiasetcontroller

import (
	"inventaris/entities"
	"inventaris/helpers"
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
	data := map[string]any{
		"Title":      "Lokasi Aset",
		"path":       path,
		"lokasiaset": lokasiaset,
		"aset":       aset,
		"bidang":     bidang,
		"pegawai":    pegawai,
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
			"Title":      "Lokasi Aset",
			"path":       path,
			"aset":       aset,
			"bidang":     bidang,
			"pegawai":    pegawai,
		}
		helpers.RenderTemplate(w, "lokasi_aset/create.html", data)
	}

	if r.Method == http.MethodPost {
		var lokasiaset entities.LokasiAset

		lokasiaset.Aset_id = r.FormValue("aset_id")
		lokasiaset.Bidang_id = r.FormValue("bidang_id")
		lokasiaset.Pegawai_id = r.FormValue("pegawai_id")
		lokasiaset.Tanggal_Perolehan, _ = time.Parse("2006-01-02", r.FormValue("tanggal_perolehan"))
		tgl_selesai, err := time.Parse("2006-01-02", r.FormValue("tanggal_selesai"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		lokasiaset.Tanggal_Selesai = &tgl_selesai
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
