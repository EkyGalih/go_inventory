package distribusicontroller

import (
	"fmt"
	"inventaris/entities"
	"inventaris/models/distribusiasetmodel"
	"inventaris/models/pegawaimodel"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Store(c *gin.Context) {
	session := sessions.Default(c)
	var statusCode int
	var bidangID string

	PegawaiID := c.PostForm("PegawaiID")
	if PegawaiID != "" {
		pegawai, _ := pegawaimodel.GetPegawaiByID(PegawaiID)
		bidID := pegawai.BidangID
		bidangID = *bidID
	} else {
		bidangID = c.PostForm("BidangID")
	}
	Ket := c.PostForm("Keterangan")
	TglPerolehan, err := time.Parse("2006-01-02", c.PostForm("TanggalPerolehan"))
	if err != nil {
		fmt.Printf("Tanggal Perolehan error, %s", err.Error())
	}
	TglSelesai, err := time.Parse("2006-01-02", c.PostForm("TanggalSelesai"))
	if err != nil {
		fmt.Printf("Tanggal Selesai error, %s", err.Error())
	}

	distribusi := entities.Distribusi{
		AsetID:           c.PostForm("AsetID"),
		PegawaiID:        &PegawaiID,
		BidangID:         bidangID,
		TanggalPerolehan: TglPerolehan,
		TanggalSelesai:   &TglSelesai,
		Keterangan:       &Ket,
	}

	if err := distribusiasetmodel.CreateDistribusi(distribusi); err != nil {
		statusCode = http.StatusInternalServerError
		session.AddFlash(fmt.Sprintf("%d: Error: %s", statusCode, err.Error()))
		session.Save()
		c.Redirect(http.StatusFound, "/addons/distribusi/"+c.PostForm("AsetID"))
		return
	}

	statusCode = http.StatusOK
	session.AddFlash(fmt.Sprintf("%d: Data berhasil disimpan", statusCode))
	session.Save()

	c.Redirect(http.StatusFound, "/aset")
}
