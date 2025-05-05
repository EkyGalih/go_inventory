package pegawaicontroller

import (
	"fmt"
	"inventaris/entities"
	"inventaris/helpers/helpers"
	"inventaris/models/bidangmodel"
	"inventaris/models/pegawaimodel"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Index(c *gin.Context) {
	// tambahkan flashes alert
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

	pegawais, _ := pegawaimodel.GetAllPegawai()

	path := map[string]string{
		"menu":    "addons",
		"subMenu": "pegawai",
	}

	data := map[string]any{
		"Title":    "Pegawai",
		"path":     path,
		"pegawais": pegawais,
		"Flashes":  messages,
	}

	helpers.RenderTemplate(c, "addons/pegawai/index.html", data)
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

	bidang, err := bidangmodel.GetAllBidang()
	if err != nil {
		session.AddFlash(fmt.Sprintf("%d: Terjadi Kesalahan: %s", http.StatusInternalServerError, err.Error()))
		session.Save()
		return
	}

	data := map[string]any{
		"Title": "Tambah Pegawai",
		"path": map[string]string{
			"menu":    "addons",
			"subMenu": "pegawai",
		},
		"bidang":  bidang,
		"Flashes": messages,
	}

	helpers.RenderTemplate(c, "addons/pegawai/create.html", data)
}

func Store(c *gin.Context) {
	// buat session
	session := sessions.Default(c)
	var statusCode int
	var publicPath string

	file, err := c.FormFile("Foto")

	if file == nil {
		publicPath = "assets/img/blank.png"
	} else {
		if err == nil {
			// buat direktori jika belum ada
			uploadDir := "public/uploads/pegawai/"
			if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
				statusCode = http.StatusInternalServerError
				session.AddFlash(fmt.Sprintf("%d: Terjadi Kesalahan: %s", statusCode, err.Error()))
				session.Save()
				c.Redirect(http.StatusFound, "/addons/pegawai/create")
				return
			}

			// generate nama file unik
			uniqueFileName := uuid.NewString() + filepath.Ext(file.Filename)
			serverFilePath := filepath.Join(uploadDir, uniqueFileName)
			publicPath = "uploads/pegawai/" + uniqueFileName

			// simpan file
			if err := c.SaveUploadedFile(file, serverFilePath); err != nil {
				statusCode = http.StatusInternalServerError
				session.AddFlash(fmt.Sprintf("%d: Terjadi Kesalahan: %s", statusCode, err.Error()))
				session.Save()
				c.Redirect(http.StatusFound, "/addons/pegawai/create")
				return
			}
		}
	}

	// buat objek pegawai
	bidangID := c.PostForm("BidangID")
	pegawai := entities.Pegawai{
		Name:         c.PostForm("Name"),
		IdPegawai:    c.PostForm("IdPegawai"),
		Foto:         publicPath,
		JenisPegawai: c.PostForm("JenisPegawai"),
		Jabatan:      c.PostForm("Jabatan"),
		BidangID:     &bidangID,
	}

	// simpan ke model
	if err := pegawaimodel.CreatePegawai(pegawai); err != nil {
		statusCode = http.StatusInternalServerError
		session.AddFlash(fmt.Sprintf("%d: Terjadi Kesalahan: %s", statusCode, err.Error()))
		session.Save()
		c.Redirect(http.StatusFound, "/addons/pegawai/create.html")
		return
	}

	statusCode = http.StatusOK
	session.AddFlash(fmt.Sprintf("%d: Data berhasil disimpan", statusCode))
	session.Save()

	c.Redirect(http.StatusFound, "/addons/pegawai")
}

func Edit(c *gin.Context) {
	// buat variabel session
	session := sessions.Default(c)
	var statusCode int
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

	id := c.Param("id")
	if id == "" {
		statusCode = http.StatusNotFound
		session.AddFlash(fmt.Sprintf("%d: Error: Id Kosong", statusCode))
		session.Save()
		c.Redirect(http.StatusFound, "/addons/pegawai/edit.html")
		return
	}

	pegawai, err := pegawaimodel.GetPegawaiByID(id)
	if err != nil {
		statusCode = http.StatusInternalServerError
		session.AddFlash(fmt.Sprintf("%d: Error: %s", statusCode, err.Error()))
		session.Save()
		c.Redirect(http.StatusFound, "/addons/pegawai/edit.html")
		return
	}

	bidang, err := bidangmodel.GetAllBidang()
	if err != nil {
		statusCode = http.StatusInternalServerError
		session.AddFlash(fmt.Sprintf("%d: Error: %s", statusCode, err.Error()))
		session.Save()
		c.Redirect(http.StatusFound, "/addons/pegawai/edit.html")
		return
	}

	BidangID := pegawai.BidangID
	var selectedBidang *entities.Bidang
	if BidangID != nil {
		selectedBidang, _ = bidangmodel.GetBidangByID(*BidangID)
	} else {
		selectedBidang = nil // Atur sesuai kebutuhan
	}

	data := map[string]any{
		"Title": "Edit Pegawai",
		"path": map[string]string{
			"menu":    "addons",
			"subMenu": "pegawai",
		},
		"pegawai":        pegawai,
		"bidang":         bidang,
		"selectedBidang": selectedBidang,
	}

	helpers.RenderTemplate(c, "addons/pegawai/edit.html", data)
}

func Update(c *gin.Context) {
	// buat variabel session
	session := sessions.Default(c)
	var statusCode int

	id := c.Param("id")
	if id == "" {
		statusCode = http.StatusNotFound
		session.AddFlash(fmt.Sprintf("%d: Error: Id Kosong", statusCode))
		session.Save()
		c.Redirect(http.StatusFound, "/addons/pegawai/edit.html")
		return
	}

	oldPegawai, err := pegawaimodel.GetPegawaiByID(id)
	if err != nil {
		statusCode = http.StatusInternalServerError
		session.AddFlash(fmt.Sprintf("%d: Error: %s", statusCode, err.Error()))
		session.Save()
		c.Redirect(http.StatusFound, "/addons/pegawai/edit.html")
		return
	}

	file, err := c.FormFile("Foto")
	publicPath := oldPegawai.Foto

	if file != nil {
		uploadDir := "public/uploads/pegawai/"
		uniqueFileName := uuid.NewString() + filepath.Ext(filepath.Base(file.Filename))
		serverFilePath := filepath.Join(uploadDir, uniqueFileName)
		publicPath = "uploads/pegawai/" + uniqueFileName

		if err := c.SaveUploadedFile(file, serverFilePath); err != nil {
			statusCode = http.StatusInternalServerError
			session.AddFlash(fmt.Sprintf("%d: Terjadi Kesalahan: %s", statusCode, err.Error()))
			session.Save()
			c.Redirect(http.StatusFound, "/addons/pegawai/edit/"+id)
			return
		}

		if oldPegawai.Foto != "assets/img/blank.png" {
			oldFilePath := "public/" + oldPegawai.Foto
			if err := os.Remove(oldFilePath); err != nil {
				session.AddFlash(fmt.Sprintf("%d: Gagal menghapus foto lama : %s", statusCode, err.Error()))
				session.Save()
				c.Redirect(http.StatusFound, "/addons/pegawai/edit/"+id)
				return
			}
		}
	}

	Name := c.PostForm("Name")
	IdPegawai := c.PostForm("IdPegawai")
	JenisPegawai := c.PostForm("JenisPegawai")
	Jabatan := c.PostForm("Jabatan")
	bidangID := c.PostForm("BidangID")

	pegawai := entities.Pegawai{
		Id:           id,
		Name:         Name,
		IdPegawai:    IdPegawai,
		JenisPegawai: JenisPegawai,
		BidangID:     &bidangID,
		Foto:         publicPath,
		Jabatan:      Jabatan,
		UpdatedAt:    time.Now(),
	}

	if err := pegawaimodel.UpdatePegawai(pegawai); err != nil {
		statusCode = http.StatusInternalServerError
		session.AddFlash(fmt.Sprintf("%d: Terjadi Kesalahan: %s", statusCode, err.Error()))
		session.Save()
		c.Redirect(http.StatusFound, "/addons/pegawai/edit/"+id)
		return
	}

	statusCode = http.StatusOK
	session.AddFlash(fmt.Sprintf("%d: Data berhasil disimpan", statusCode))
	session.Save()
	c.Redirect(http.StatusFound, "/addons/pegawai")
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
		c.Redirect(http.StatusFound, "/addons/pegawai")
		return
	}
	pegawai, err := pegawaimodel.GetPegawaiByID(id)
	if err != nil {
		statusCode = http.StatusInternalServerError
		session.AddFlash(fmt.Sprintf("%d: Error mengambil data pegawai : %s", statusCode, err.Error()))
		session.Save()
		c.Redirect(http.StatusFound, "/addons/pegawai")
		return
	}

	// jika foto bukan default foto, maka hapus
	if pegawai.Foto != "assets/img/blank.png" {
		filePath := "public/" + pegawai.Foto
		if err := os.Remove(filePath); err != nil {
			session.AddFlash(fmt.Sprintf("%d: Gagal menghapus foto pegawai: %s", http.StatusInternalServerError, err.Error()))
		}
	}

	err = pegawaimodel.DeletePegawai(id)
	if err != nil {
		statusCode := http.StatusInternalServerError
		session.AddFlash(fmt.Sprintf("%d: Gagal menghapus data pegawai, masalah : %s", statusCode, err.Error()))
	}

	statusCode = http.StatusOK
	session.AddFlash(fmt.Sprintf("%d: Data berhasil dihapus", statusCode))
	session.Save()
	c.Redirect(http.StatusFound, "/addons/pegawai")
}
