package lokasiasetcontroller

import (
	"encoding/json"
	"inventaris/entities"
	"inventaris/helpers/helpers"
	"inventaris/helpers/queryhelpers"
	"inventaris/models/asettikmodel"
	"inventaris/models/bidangmodel"
	"inventaris/models/lokasiasetmodel"
	"inventaris/models/pegawaimodel"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
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
		aset := asettikmodel.GetDataAset()
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

		aset_id := r.FormValue("aset_id")
		bidang_id := r.FormValue("bidang_id")
		pegawai_id := r.FormValue("pegawai_id")

		aset, err := asettikmodel.Detail(aset_id)
		if err != nil {
			http.Error(w, "failed to get detail aset : "+err.Error(), http.StatusInternalServerError)
			return
		}

		bidang, err := bidangmodel.Detail(bidang_id)
		if err != nil {
			http.Error(w, "failed to get detail bidang : "+err.Error(), http.StatusInternalServerError)
			return
		}

		pegawai, err := pegawaimodel.Detail(pegawai_id)
		if err != nil {
			http.Error(w, "failed to get detail pegawai : "+err.Error(), http.StatusInternalServerError)
			return
		}
		
		lokasiaset.Aset_id = r.FormValue("aset_id")
		lokasiaset.Bidang_id = r.FormValue("bidang_id")
		lokasiaset.Pegawai_id = r.FormValue("pegawai_id")
		tanggalPerolehan, err := time.Parse("2006-01-02", r.FormValue("tanggal_perolehan"))
		if err != nil {
			http.Error(w, "Invalid date format for Tanggal Perolehan", http.StatusBadRequest)
			return
		}
		lokasiaset.Tanggal_Perolehan = tanggalPerolehan
		var tanggalSelesai *time.Time
		tglSelesaiStr := r.FormValue("tanggal_selesai")
		if tglSelesaiStr != "" {
			tglSelesai, err := time.Parse("2006-01-02", tglSelesaiStr)
			if err != nil {
				http.Error(w, "Invalid date format for Tanggal Selesai", http.StatusBadRequest)
				return
			}
			tanggalSelesai = &tglSelesai
		}
		lokasiaset.Tanggal_Selesai = tanggalSelesai
		jenis_pemanfaatan := r.FormValue("jenis_pemanfaatan")
		lokasiaset.Jenis_Pemanfaatan = &jenis_pemanfaatan
		keterangan := r.FormValue("keterangan")
		lokasiaset.Keterangan = &keterangan
		lokasiaset.Created_At = time.Now()
		lokasiaset.Updated_At = time.Now()

		data := entities.Riwayat{
			Id:           uuid.New().String(),
			Aset_id:      aset.Id,
			Nama_Aset:    aset.Nama_Aset,
			Kode_Aset:    aset.Kode_Aset,
			Bidang_id:    bidang.Id,
			Nama_Bidang:  bidang.Nama_Bidang,
			Pegawai_id:   pegawai.Id,
			Nama_Pegawai: pegawai.Name,
			Foto_Pegawai: pegawai.Foto,
			Nip_Pegawai:  pegawai.Nip,
			Tanggal_Aksi: tanggalPerolehan,
			Jenis_Aksi:   "Penerimaan Aset",
			Keterangan:   &keterangan,
			Created_At:   time.Now(),
			Updated_At:   time.Now(),
		}

		var existingData []entities.Riwayat
		path := "./data/riwayataset"
		jsonName := aset.Kode_Aset + ".json"
		jsonFile := filepath.Join(path, jsonName)

		// Membaca data JSON yang ada
		if _, err := os.Stat(jsonFile); !os.IsNotExist(err) {
			file, err := os.ReadFile(jsonFile)
			if err != nil {
				http.Error(w, "Failed to read existing JSON file : "+err.Error(), http.StatusInternalServerError)
				return
			}

			err = json.Unmarshal(file, &existingData)
			if err != nil {
				http.Error(w, "Failed to parse existing JSON data : "+err.Error(), http.StatusInternalServerError)
				return
			}
		}

		// Menambahkan data baru ke array
		existingData = append(existingData, data)

		jsonData, err := json.MarshalIndent(existingData, "", "   ")
		if err != nil {
			http.Error(w, "Failed to Marshal data json :"+err.Error(), http.StatusInternalServerError)
			return
		}

		err = os.WriteFile(jsonFile, jsonData, 0644)
		if err != nil {
			http.Error(w, "Failed to write JSON to file", http.StatusInternalServerError)
			return
		}

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
