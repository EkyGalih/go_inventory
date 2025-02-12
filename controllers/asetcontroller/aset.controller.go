package asetcontroller

import (
	"fmt"
	"inventaris/entities"
	"inventaris/helpers/helpers"
	"inventaris/helpers/queryhelpers"
	"inventaris/models/asetmodel"
	"inventaris/models/bidangmodel"
	"inventaris/models/categorymodel"
	"inventaris/models/pegawaimodel"
	"inventaris/models/tipemodel"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Index(c *gin.Context) {
	// tambahkan flashses alert
	session := sessions.Default(c)
	flashes := session.Flashes()
	session.Save()
	var message []map[string]interface{}

	for _, flash := range flashes {
		parts := strings.SplitN(flash.(string), ": ", 2)
		if len(parts) == 2 {
			status, _ := strconv.Atoi(parts[0])
			message = append(message, map[string]interface{}{
				"status":  status,
				"message": parts[1],
			})
		}
	}

	asets, _ := asetmodel.GetAllAset()

	distribusi := queryhelpers.GetDistribusi(asets)
	if distribusi == nil {
		distribusi = make(map[string]int)
	}

	pemegangAsets := queryhelpers.GetPemegangAsets(asets)
	if pemegangAsets == nil {
		pemegangAsets = make(map[string][]entities.Distribusi)
	}

	data := map[string]interface{}{
		"Title": "Aset Tetap",
		"path": map[string]string{
			"menu":    "aset",
			"subMenu": "aset-tetap",
		},
		"asets":         asets,
		"Distribusi":    distribusi,
		"PemegangAsets": pemegangAsets,
		"Flashes":       message,
	}

	helpers.RenderTemplate(c, "aset/index.html", data)
}

func Add(c *gin.Context) {
	session := sessions.Default(c)
	flashes := session.Flashes()
	session.Save()
	var messages []map[string]interface{}

	for _, flash := range flashes {
		parts := strings.SplitN(flash.(string), ": ", 2)
		if len(parts) == 2 {
			status, _ := strconv.Atoi(parts[0])
			messages = append(messages, map[string]interface{}{
				"status":  status,
				"message": parts[1],
			})
		}
	}

	kategori, err := categorymodel.GetAllCategory()
	if err != nil {
		session.AddFlash(fmt.Sprintf("%d: Terjadi Kesalahan : %s", http.StatusInternalServerError, err.Error()))
		session.Save()
		c.Redirect(http.StatusFound, "aset/create")
		return
	}

	tipe, err := tipemodel.GetAllTipe()
	if err != nil {
		session.AddFlash(fmt.Sprintf("%d: Terjadi Kesalahan : %s", http.StatusInternalServerError, err.Error()))
		session.Save()
		c.Redirect(http.StatusFound, "aset/create")
		return
	}
	data := map[string]any{
		"Title": "Tambah Aset Tetap",
		"path": map[string]string{
			"menu":    "aset",
			"subMenu": "aset-tetap",
		},
		"Flashes":    flashes,
		"tipes":      tipe,
		"categories": kategori,
	}

	helpers.RenderTemplate(c, "aset/create.html", data)
}

func Store(c *gin.Context) {
	// buat session
	session := sessions.Default(c)
	var statusCode int
	var publicPath string

	file, err := c.FormFile("Gambar")

	if file == nil {
		publicPath = "asset/img/image.png"
	} else {
		if err == nil {
			// buat direktori jika belum ada
			uploadDir := "public/uploads/aset/"
			if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
				statusCode = http.StatusInternalServerError
				session.AddFlash(fmt.Sprintf("%d : Terjadi kesalahan: %s", statusCode, err.Error()))
				session.Save()
				c.Redirect(http.StatusFound, "/aset/create")
				return
			}

			// generate nama file unik
			uniqueFileName := uuid.NewString() + filepath.Ext(file.Filename)
			serverFilePath := filepath.Join(uploadDir, uniqueFileName)
			publicPath = "uploads/aset/" + uniqueFileName

			// simpan file
			if err := c.SaveUploadedFile(file, serverFilePath); err != nil {
				statusCode = http.StatusInternalServerError
				session.AddFlash(fmt.Sprintf("%d: Terjadi Kesalahan: %s", statusCode, err.Error()))
				session.Save()
				c.Redirect(http.StatusFound, "/aset/create")
				return
			}
		}
	}
	// buat objek aset
	aset := entities.Aset{
		JenisAset:    c.PostForm("JenisAset"),
		KodeAset:     c.PostForm("KodeAset"),
		NamaAset:     c.PostForm("NamaAset"),
		Merek:        c.PostForm("Merek"),
		Model:        c.PostForm("Model"),
		SerialNumber: c.PostForm("SerialNumber"),
		Deskripsi:    c.PostForm("Deskripsi"),
		KategoriID:   c.PostForm("KategoriID"),
		TipeID:       c.PostForm("TipeID"),
		TanggalPerolehan: func() time.Time {
			tanggalPerolehan, _ := time.Parse("2006-01-02", c.PostForm("TanggalPerolehan"))
			return tanggalPerolehan
		}(),
		Status: c.PostForm("Status"),
		Nilai: func() float64 {
			input := c.PostForm("Nilai")
			re := regexp.MustCompile(`[^0-9,.]`)
			cleaned := re.ReplaceAllString(input, "")

			cleaned = strings.ReplaceAll(cleaned, ",", ".")

			nilai, _ := strconv.ParseFloat(cleaned, 64)
			return nilai
		}(),
		Jumlah: func() int {
			jumlah, _ := strconv.Atoi(c.PostForm("Jumlah"))
			return jumlah
		}(),
		Keterangan: c.PostForm("Keterangan"),
		Path:       publicPath,
		Satuan:     c.PostForm("Satuan"),
	}

	// simpan ke model
	if err := asetmodel.CreateAset(aset); err != nil {
		statusCode = http.StatusInternalServerError
		session.AddFlash(fmt.Sprintf("%d: Terjadi Kesalahan: %s", statusCode, err.Error()))
		session.Save()
		c.Redirect(http.StatusFound, "/aset/create.html")
		return
	}

	statusCode = http.StatusOK
	session.AddFlash(fmt.Sprintf("%d: Data berhasil disimpan", statusCode))
	session.Save()

	c.Redirect(http.StatusFound, "/aset")
}

func Edit(c *gin.Context) {
	session := sessions.Default(c)
	flashes := session.Flashes()
	session.Save()
	var statusCode int
	var messages []map[string]interface{}

	for _, flash := range flashes {
		parts := strings.SplitN(flash.(string), ": ", 2)
		if len(parts) == 2 {
			status, _ := strconv.Atoi(parts[0])
			messages = append(messages, map[string]interface{}{
				"status":  status,
				"message": parts[1],
			})
		}
	}

	id := c.Param("id")
	if id == "" {
		statusCode = http.StatusNotFound
		session.AddFlash(fmt.Sprintf("%d: Error: Id Kosong", statusCode))
		session.Save()
		c.Redirect(http.StatusFound, "/aset/edit.html")
		return
	}

	aset, err := asetmodel.GetAsetByID(id)
	if err != nil {
		statusCode = http.StatusInternalServerError
		session.AddFlash(fmt.Sprintf("%d: Error: %s", statusCode, err.Error()))
		session.Save()
		c.Redirect(http.StatusFound, "/aset/edit.html")
		return
	}

	kategori, err := categorymodel.GetAllCategory()
	if err != nil {
		session.AddFlash(fmt.Sprintf("%d: Terjadi Kesalahan : %s", http.StatusInternalServerError, err.Error()))
		session.Save()
		c.Redirect(http.StatusFound, "aset/edit")
		return
	}

	tipe, err := tipemodel.GetAllTipe()
	if err != nil {
		session.AddFlash(fmt.Sprintf("%d: Terjadi Kesalahan : %s", http.StatusInternalServerError, err.Error()))
		session.Save()
		c.Redirect(http.StatusFound, "aset/edit")
		return
	}

	KategoriID, _ := categorymodel.GetCategoryByID(aset.KategoriID)
	TipeID, _ := tipemodel.GetTipeByID(aset.TipeID)

	data := map[string]any{
		"Title": "Tambah Aset Tetap",
		"path": map[string]string{
			"menu":    "aset",
			"subMenu": "aset-tetap",
		},
		"Flashes":          flashes,
		"Aset":             aset,
		"tipe":             tipe,
		"SelectedTipe":     TipeID.ID,
		"kategori":         kategori,
		"SelectedKategori": KategoriID.ID,
	}

	helpers.RenderTemplate(c, "aset/edit.html", data)
}

func Update(c *gin.Context) {
	// buat session
	session := sessions.Default(c)
	var statusCode int

	id := c.Param("id")
	if id == "" {
		statusCode = http.StatusNotFound
		session.AddFlash(fmt.Sprintf("%d: Error: Id Kosong", statusCode))
		session.Save()
		c.Redirect(http.StatusFound, "/aset/edit.html")
		return
	}

	oldAset, err := asetmodel.GetAsetByID(id)
	if err != nil {
		statusCode = http.StatusNotFound
		session.AddFlash(fmt.Sprintf("%d: Error: %s", statusCode, err.Error()))
		session.Save()
		c.Redirect(http.StatusFound, "/aset/edit.html")
		return
	}

	file, err := c.FormFile("Gambar")
	publicPath := oldAset.Path

	if file != nil {
		// buat direktori jika belum ada
		uploadDir := "public/uploads/aset/"
		uniqueFileName := uuid.NewString() + filepath.Ext(filepath.Base(file.Filename))
		serverFilePath := filepath.Join(uploadDir, uniqueFileName)
		publicPath = "uploads/aset/" + uniqueFileName

		// simpan file
		if err := c.SaveUploadedFile(file, serverFilePath); err != nil {
			statusCode = http.StatusInternalServerError
			session.AddFlash(fmt.Sprintf("%d: Terjadi Kesalahan saat mengupload file: %s", statusCode, err.Error()))
			session.Save()
			c.Redirect(http.StatusFound, "/aset/edit/"+id)
			return
		}

		if oldAset.Path != "assets/img/image.png" {
			oldFilePath := "public/" + oldAset.Path
			if err := os.Remove(oldFilePath); err != nil {
				session.AddFlash(fmt.Sprintf("%d: Gagal menghapus foto lama : %s", statusCode, err.Error()))
				session.Save()
				c.Redirect(http.StatusFound, "/aset/edit/"+id)
				return
			}
		}
	}

	// buat objek aset
	aset := entities.Aset{
		ID:           id,
		JenisAset:    c.PostForm("JenisAset"),
		KodeAset:     c.PostForm("KodeAset"),
		NamaAset:     c.PostForm("NamaAset"),
		Merek:        c.PostForm("Merek"),
		Model:        c.PostForm("Model"),
		SerialNumber: c.PostForm("SerialNumber"),
		Deskripsi:    c.PostForm("Deskripsi"),
		KategoriID:   c.PostForm("KategoriID"),
		TipeID:       c.PostForm("TipeID"),
		TanggalPerolehan: func() time.Time {
			tanggalPerolehan, _ := time.Parse("2006-01-02", c.PostForm("TanggalPerolehan"))
			return tanggalPerolehan
		}(),
		Status: c.PostForm("Status"),
		Nilai: func() float64 {
			input := c.PostForm("Nilai")
			re := regexp.MustCompile(`[^0-9,.]`)
			cleaned := re.ReplaceAllString(input, "")

			cleaned = strings.ReplaceAll(cleaned, ".", "")
			cleaned = strings.ReplaceAll(cleaned, ",", ".")

			nilai, err := strconv.ParseFloat(cleaned, 64)
			if err != nil {
				fmt.Println("Error parsing float:", err) // Debugging
				return 0
			}
			return nilai
		}(),
		Jumlah: func() int {
			jumlah, _ := strconv.Atoi(c.PostForm("Jumlah"))
			return jumlah
		}(),
		Keterangan: c.PostForm("Keterangan"),
		Path:       publicPath,
		Satuan:     c.PostForm("Satuan"),
		UpdatedAt:  time.Now(),
	}

	// simpan ke model
	if err := asetmodel.UpdateAset(aset); err != nil {
		statusCode = http.StatusInternalServerError
		session.AddFlash(fmt.Sprintf("%d: Terjadi Kesalahan saat update aset: %s", statusCode, err.Error()))
		session.Save()
		c.Redirect(http.StatusFound, "/aset/edit/"+id)
		return
	}

	statusCode = http.StatusOK
	session.AddFlash(fmt.Sprintf("%d: Data berhasil disimpan", statusCode))
	session.Save()

	c.Redirect(http.StatusFound, "/aset")
}

func Distribusi(c *gin.Context) {
	// buat session
	session := sessions.Default(c)
	flashes := session.Flashes()
	session.Save()

	idAset := c.Param("id")
	aset, _ := asetmodel.GetAsetByID(idAset)
	pegawai, _ := pegawaimodel.GetAllPegawai()
	bidang, _ := bidangmodel.GetAllBidang()

	data := map[string]interface{}{
		"Title":   "Distribusi Aset",
		"path":    map[string]string{"menu": "aset"},
		"Flashes": flashes,
		"aset":    aset,
		"pegawai": pegawai,
		"bidang":  bidang,
	}

	helpers.RenderTemplate(c, "/aset/distribusi.html", data)
}

func Delete(c *gin.Context) {
	// buat variabel session dan statusCode
	session := sessions.Default(c)
	var statusCode int

	id := c.Param("id")
	if id == "" {
		statusCode = http.StatusNotFound
		session.AddFlash(fmt.Sprintf("%d: Error: Id Kosong", statusCode))
		session.Save()
		c.Redirect(http.StatusFound, "/aset/tetap")
		return
	}
	aset, err := asetmodel.GetAsetByID(id)
	if err != nil {
		statusCode = http.StatusInternalServerError
		session.AddFlash(fmt.Sprintf("%d: Error mengambil data aset : %s", statusCode, err.Error()))
		session.Save()
		c.Redirect(http.StatusFound, "/aset")
		return
	}

	// jika foto bukan default foto, maka hapus
	if aset.Path != "assets/img/image.png" {
		filePath := "public/" + aset.Path
		if err := os.Remove(filePath); err != nil {
			session.AddFlash(fmt.Sprintf("%d: Gagal menghapus foto aset: %s", http.StatusInternalServerError, err.Error()))
		}
	}

	err = asetmodel.DeleteAset(id)
	if err != nil {
		statusCode := http.StatusInternalServerError
		session.AddFlash(fmt.Sprintf("%d: Gagal menghapus data aset, masalah : %s", statusCode, err.Error()))
	}

	statusCode = http.StatusOK
	session.AddFlash(fmt.Sprintf("%d: Data berhasil dihapus", statusCode))
	session.Save()
	c.Redirect(http.StatusFound, "/aset")
}
