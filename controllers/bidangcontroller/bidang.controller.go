package bidangcontroller

import (
	"inventaris/helpers/helpers"
	"inventaris/models/bidangmodel"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	// tambahkan flashes alert
	session := sessions.Default(c)
	flashes := session.Flashes()
	session.Save()

	bidangs, _ := bidangmodel.GetAllBidang()
	path := map[string]string{
		"menu":    "addons",
		"subMenu": "bidang",
	}

	data := map[string]any{
		"Title":   "Bidang",
		"path":    path,
		"bidangs": bidangs,
		"Flashes": flashes,
	}

	helpers.RenderTemplate(c, "addons/bidang/index.html", data)
}

func AddBidang(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		data := map[string]any{
			"Title": "Tambah Bidang",
			"path": map[string]string{
				"menu":    "addons",
				"subMenu": "bidang",
			},
		}

		helpers.RenderTemplate(c, "/addons/bidang/create.html", data)
	}

	if c.Request.Method == http.MethodPost {
		namaBidang := c.PostForm("NamaBidang")
		bidangmodel.CreateBidang(namaBidang)
		session := sessions.Default(c)
		session.AddFlash("Data berhasil disimpan")
		session.Save()

		c.Redirect(http.StatusFound, "/addons/bidang")
	}
}

func EditBidang(c *gin.Context) {
	session := sessions.Default(c)

	if c.Request.Method == http.MethodGet {
		id := c.Param("id")
		if id == "" {
			session.AddFlash("Id Kosong")
			session.Save()
		}
		bidang, err := bidangmodel.GetBidangByID(id)
		if err != nil {
			session.AddFlash("Bidang tidak ditemukan")
			session.Save()
		}

		data := map[string]interface{}{
			"Title": "Edit Bidang",
			"path": map[string]string{
				"menu":    "addons",
				"subMenu": "bidang",
			},
			"bidang": bidang,
		}

		helpers.RenderTemplate(c, "/addons/bidang/edit.html", data)
	}

	if c.Request.Method == http.MethodPost {
		id := c.Param("id")
		log.Println(id)
		namaBidang := c.PostForm("NamaBidang")
		bidangmodel.UpdateBidang(id, namaBidang)
		session.AddFlash("Data berhasil disimpan")
		session.Save()
		c.Redirect(http.StatusFound, "/addons/bidang")
	}
}

func DeleteBidang(c *gin.Context) {
	id := c.Param("id")
	bidangmodel.DeleteBidang(id)
	session := sessions.Default(c)
	session.AddFlash("Data berhasil dihapus")
	session.Save()
	c.Redirect(http.StatusFound, "/addons/bidang")
}
