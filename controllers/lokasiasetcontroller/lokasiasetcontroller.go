package lokasiasetcontroller

import (
	"database/sql"
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
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Index(c *gin.Context) {
	// buat session message
	session := sessions.Default(c)
	flashes := session.Flashes()
	session.Save()

	// Mengambil semua lokasi aset, aset, bidang, dan pegawai
	lokasiaset, err := lokasiasetmodel.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve lokasi aset: " + err.Error()})
		return
	}
	aset, _ := asettikmodel.GetAll()
	bidang, _ := bidangmodel.GetAllBidang()
	pegawai, _ := pegawaimodel.GetAll()

	// Menghitung jumlah aset per pegawai
	aset_pegawai := queryhelpers.CountAsetPegawai(lokasiaset)
	if aset_pegawai == nil {
		aset_pegawai = make(map[string]int)
	}

	// Data untuk template
	data := map[string]interface{}{
		"Title":        "Lokasi Aset",
		"path":         map[string]string{"menu": "lokasi-aset"},
		"lokasiaset":   lokasiaset,
		"aset":         aset,
		"bidang":       bidang,
		"pegawai":      pegawai,
		"aset_pegawai": aset_pegawai,
		"flashes":      flashes,
	}

	helpers.RenderTemplate(c, "lokasi_aset/index.html", data)
}

func Add(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		// Mengambil data untuk form input
		aset, _ := asettikmodel.GetAll()
		bidang, _ := bidangmodel.GetAllBidang()
		pegawai, _ := pegawaimodel.GetAll()

		data := gin.H{
			"Title":   "Lokasi Aset",
			"path":    map[string]string{"menu": "lokasi-aset"},
			"aset":    aset,
			"bidang":  bidang,
			"pegawai": pegawai,
		}
		if err := helpers.RenderTemplate(c, "lokasi_aset/create.html", data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	if c.Request.Method == http.MethodPost {
		var lokasiaset entities.LokasiAset

		// Mengambil nilai dari form input
		aset_id := c.PostForm("aset_id")
		bidang_id := c.PostForm("bidang_id")
		pegawai_id := c.PostForm("pegawai_id")

		// Mengambil data terkait berdasarkan ID
		aset, err := asettikmodel.Detail(aset_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get detail aset: " + err.Error()})
			return
		}

		bidang, err := bidangmodel.GetBidangByID(bidang_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get detail bidang: " + err.Error()})
			return
		}

		pegawai, err := pegawaimodel.Detail(pegawai_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get detail pegawai: " + err.Error()})
			return
		}

		// Mengisi data lokasiaset dari form
		lokasiaset.AsetID = aset_id
		lokasiaset.BidangID = bidang_id
		lokasiaset.PegawaiID = pegawai_id
		tanggalPerolehan, err := time.Parse("2006-01-02", c.PostForm("tanggal_perolehan"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format for Tanggal Perolehan"})
			return
		}
		lokasiaset.TanggalPerolehan = tanggalPerolehan

		// Menangani tanggal selesai
		var tanggalSelesai *time.Time
		tglSelesaiStr := c.PostForm("tanggal_selesai")
		if tglSelesaiStr != "" {
			tglSelesai, err := time.Parse("2006-01-02", tglSelesaiStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format for Tanggal Selesai"})
				return
			}
			tanggalSelesai = &tglSelesai
		}
		lokasiaset.TanggalSelesai = tanggalSelesai
		jenisPemanfaatan := c.PostForm("jenis_pemanfaatan")
		lokasiaset.JenisPemanfaatan = &jenisPemanfaatan
		keterangan := c.PostForm("keterangan")
		lokasiaset.Keterangan = &keterangan
		lokasiaset.CreatedAt = time.Now()
		lokasiaset.UpdatedAt = time.Now()

		// Membuat riwayat aset dan menyimpannya ke file JSON
		data := entities.Riwayat{
			ID:          uuid.New().String(),
			AsetID:      aset.ID,
			NamaAset:    aset.NamaAset,
			KodeAset:    aset.KodeAset,
			BidangID:    bidang.ID,
			NamaBidang:  bidang.NamaBidang,
			PegawaiID:   strconv.Itoa(int(pegawai.ID)),
			NamaPegawai: pegawai.Name,
			FotoPegawai: sql.NullString{String: pegawai.Foto, Valid: pegawai.Foto != ""}, // Convert to sql.NullString
			NipPegawai:  sql.NullString{String: pegawai.Nip, Valid: pegawai.Nip != ""},   // Convert to sql.NullString
			TanggalAksi: tanggalPerolehan,
			JenisAksi:   "Penerimaan Aset",
			Keterangan:  &keterangan,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		// Menyimpan riwayat ke file JSON
		path := "./data/riwayataset"
		jsonName := aset.KodeAset + ".json"
		jsonFile := filepath.Join(path, jsonName)

		var existingData []entities.Riwayat
		if _, err := os.Stat(jsonFile); !os.IsNotExist(err) {
			file, err := os.ReadFile(jsonFile)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read existing JSON file: " + err.Error()})
				return
			}

			err = json.Unmarshal(file, &existingData)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse existing JSON data: " + err.Error()})
				return
			}
		}

		existingData = append(existingData, data)
		jsonData, err := json.MarshalIndent(existingData, "", "   ")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Marshal data json: " + err.Error()})
			return
		}

		err = os.WriteFile(jsonFile, jsonData, 0644)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write JSON to file"})
			return
		}

		// Menyimpan data lokasi aset menggunakan GORM
		success, err := lokasiasetmodel.Create(lokasiaset)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if success {
			c.Redirect(http.StatusSeeOther, "/lokasi-aset")
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create lokasi aset"})
		}
	}
}

func Edit(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		idString := c.Query("id")
		if idString == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter id tidak ditemukan"})
			return
		}

		// Mengambil data untuk form edit
		aset, _ := asettikmodel.GetAll()
		bidang, _ := bidangmodel.GetAllBidang()
		pegawai, _ := pegawaimodel.GetAll()

		// Mengambil detail lokasi aset
		lokasiaset, err := lokasiasetmodel.Detail(idString)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		data := map[string]interface{}{
			"Title":             "Edit Lokasi Aset",
			"path":              map[string]string{"menu": "lokasi-aset"},
			"aset":              aset,
			"bidang":            bidang,
			"pegawai":           pegawai,
			"lokasiaset":        lokasiaset,
			"SelectedAset":      lokasiaset.AsetID,
			"SelectedBidang":    lokasiaset.BidangID,
			"SelectedPegawai":   lokasiaset.PegawaiID,
			"SelectedJenisAset": lokasiaset.JenisPemanfaatan,
		}

		c.HTML(http.StatusOK, "lokasi_aset/edit.html", data)
	}

	if c.Request.Method == http.MethodPost {
		// Memperbarui data lokasi aset
		idString := c.PostForm("id")
		if idString == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter id tidak ditemukan"})
			return
		}

		var lokasiaset entities.LokasiAset
		lokasiaset.AsetID = c.PostForm("aset_id")
		lokasiaset.BidangID = c.PostForm("bidang_id")
		lokasiaset.PegawaiID = c.PostForm("pegawai_id")
		lokasiaset.TanggalPerolehan, _ = time.Parse("2006-01-02", c.PostForm("tanggal_perolehan"))
		tglSelesaiStr := c.PostForm("tanggal_selesai")
		if tglSelesaiStr != "" {
			tgl_selesai, err := time.Parse("2006-01-02", tglSelesaiStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			lokasiaset.TanggalSelesai = &tgl_selesai
		} else {
			lokasiaset.TanggalSelesai = nil
		}
		jenisPemanfaatan := c.PostForm("jenis_pemanfaatan")
		lokasiaset.JenisPemanfaatan = &jenisPemanfaatan
		keterangan := c.PostForm("keterangan")
		lokasiaset.Keterangan = &keterangan
		lokasiaset.UpdatedAt = time.Now()

		success, err := lokasiasetmodel.Update(idString, lokasiaset)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if success {
			c.Redirect(http.StatusSeeOther, "/lokasi-aset")
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update lokasi aset"})
		}
	}
}
