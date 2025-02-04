package tipeasetcontroller

import (
	"fmt"
	"inventaris/entities"
	"inventaris/helpers/helpers"
	"inventaris/models/tipemodel"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	// tambahkan flases alert
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

	tipes, _ := tipemodel.GetAllTipe()
	
	data := map[string]any{
		"Title":   "Tipe Aset",
		"path": map[string]string{
			"menu":    "addons",
			"subMenu": "type",
		},
		"tipes": tipes,
		"Flashes": messages,
	}

	helpers.RenderTemplate(c, "/addons/type/index.html", data)
}

func Add(c *gin.Context) {
	// buat session
	session := sessions.Default(c)
	var statusCode int
	
	ket := c.PostForm("Keterangan")
	tipe := entities.Tipe{
		Nama_Tipe: c.PostForm("Nama_Tipe"),
		Keterangan: &ket,
	}
	
	if err := tipemodel.CreateTipe(tipe); err != nil {
		statusCode = http.StatusInternalServerError
		session.AddFlash(fmt.Sprintf("%d: Terjadi Kesalahan: %s", statusCode, err.Error()))
		session.Save()
		c.Redirect(http.StatusFound, "/addons/type")
		return
	}

	statusCode = http.StatusOK
	session.AddFlash(fmt.Sprintf("%d: Data berhasil disimpan", statusCode))
	session.Save()
	c.Redirect(http.StatusFound, "/addons/tipe")
}

func Edit(c *gin.Context) {
	// buat variabel session
	session := sessions.Default(c)
	var statusCode int

	id := c.Param("id")
	if id == "" {
		statusCode = http.StatusNotFound
		session.AddFlash(fmt.Sprintf("%d: Error: Id Kosong", statusCode))
		session.Save()
		c.Redirect(http.StatusFound, "/addons/tipe")
		return
	}

	var tipe entities.Tipe
	if err := c.ShouldBind(&tipe); err != nil {
		statusCode = http.StatusBadRequest
		session.AddFlash(fmt.Sprintf("%d: Error: %s", statusCode, err.Error()))
		session.Save()
		c.Redirect(http.StatusFound, "/addons/tipe")
		return
	}

	tipe.ID = id
	tipe.UpdatedAt = time.Now()

	err := tipemodel.UpdateTipe(tipe)
	if err != nil {
		statusCode = http.StatusInternalServerError
		session.AddFlash(fmt.Sprintf("%d: Error: %s", statusCode, err.Error()))
		session.Save()
		c.Redirect(http.StatusFound, "/addons/tipe")
		return
	}

	statusCode = http.StatusOK
	session.AddFlash(fmt.Sprintf("%d: Data berhasil disimpan", statusCode))
	session.Save()
	c.Redirect(http.StatusFound, "/addons/tipe")
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
		c.Redirect(http.StatusFound, "/addons/tipe")
		return
	}
	err := tipemodel.DeleteTipe(id)
	if err != nil {
		statusCode = http.StatusInternalServerError
		session.AddFlash(fmt.Sprintf("%d: Error: %s", statusCode, err.Error()))
		session.Save()
		c.Redirect(http.StatusFound, "/addons/tipe")
		return
	}

	statusCode = http.StatusOK
	session.AddFlash(fmt.Sprintf("%d: Data berhasil dihapus", statusCode))
	session.Save()
	c.Redirect(http.StatusFound, "/addons/tipe")
}