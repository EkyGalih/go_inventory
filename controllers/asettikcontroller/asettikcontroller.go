package asettikcontroller

// import (
// 	"inventaris/entities"
// 	"inventaris/helpers/helpers"
// 	"inventaris/helpers/queryhelpers"
// 	"inventaris/models/asettikmodel"
// 	"inventaris/models/bidangmodel"
// 	"inventaris/models/categorymodel"
// 	"inventaris/models/pegawaimodel"
// 	"inventaris/models/tipemodel"
// 	"io"
// 	"math"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// func Index(c *gin.Context) {
// 	pageStr := c.Query("page")
// 	limitStr := c.Query("limit")

// 	page, err := strconv.Atoi(pageStr)
// 	if err != nil || page < 1 {
// 		page = 1
// 	}

// 	limit, err := strconv.Atoi(limitStr)
// 	if err != nil || limit < 1 {
// 		limit = 10
// 	}

// 	aset_tiks, err := asettikmodel.GetPaginate(page, limit)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	aset_tik, err := asettikmodel.GetAll()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	totalRows, err := asettikmodel.GetTotalRows()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	totalPages := int(math.Ceil(float64(totalRows) / float64(limit)))

// 	distribusi := queryhelpers.GetDistribusi(aset_tik)
// 	if distribusi == nil {
// 		distribusi = make(map[string]int)
// 	}

// 	// untuk show menu dan active sub menu
// 	path := map[string]string{
// 		"menu":    "aset",
// 		"subMenu": "aset-tik",
// 	}

// 	data := map[string]any{
// 		"Title":      "Aset TIK",
// 		"path":       path,
// 		"Page":       page,
// 		"TotalPages": totalPages,
// 		"TotalRows":  totalRows,
// 		"Limit":      limit,
// 		"aset_tiks":  aset_tiks,
// 		"distribusi": distribusi,
// 	}

// 	c.HTML(http.StatusOK, "/aset/aset_tik/index.html", data)
// }

// func Add(c *gin.Context) {
// 	if c.Request.Method == http.MethodGet {
// 		categories, err := categorymodel.GetAll()
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		tipes, err := tipemodel.GetAll()
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		path := map[string]string{
// 			"menu":    "aset",
// 			"subMenu": "aset-tik",
// 		}

// 		data := map[string]any{
// 			"Title":      "Add Aset TIK",
// 			"path":       path,
// 			"categories": categories,
// 			"tipes":      tipes,
// 		}
// 		c.HTML(http.StatusOK, "/aset/aset_tik/create.html", data)
// 	}

// 	if c.Request.Method == http.MethodPost {
// 		err := c.Request.ParseMultipartForm(10 << 20)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		var aset_tik entities.AsetTik

// 		file, _, err := c.Request.FormFile("Gambar")
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to retrieve file " + err.Error()})
// 			return
// 		}
// 		defer file.Close()

// 		path := "./public/uploads/aset/tetap/"
// 		fileName := helpers.RandString(30) + ".jpg"
// 		filePath := filepath.Join(path, fileName)
// 		dbPath := strings.ReplaceAll(filepath.Join("/public/uploads/aset/tetap/", fileName), "\\", "/")

// 		dest, err := os.Create(filePath)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file: " + err.Error()})
// 			return
// 		}
// 		defer dest.Close()
// 		_, err = io.Copy(dest, file)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file: " + err.Error()})
// 			return
// 		}

// 		aset_tik.JenisAset = c.PostForm("JenisAset")
// 		aset_tik.KodeAset = c.PostForm("KodeAset")
// 		aset_tik.NamaAset = c.PostForm("NamaAset")
// 		aset_tik.Merek = c.PostForm("Merek")
// 		aset_tik.Model = c.PostForm("Model")
// 		aset_tik.SerialNumber = c.PostForm("SerialNumber")
// 		deskripsi := c.PostForm("Deskripsi")
// 		aset_tik.Deskripsi = deskripsi
// 		aset_tik.KategoriID = c.PostForm("KategoriID")
// 		aset_tik.TipeID = c.PostForm("TipeID")
// 		aset_tik.TanggalPerolehan, _ = time.Parse("2006-01-02", c.PostForm("TanggalPerolehan"))
// 		aset_tik.Status = c.PostForm("Status")
// 		aset_tik.Nilai, _ = helpers.ParseCurrencyToFloat(c.PostForm("Nilai"))
// 		jumlahFloat, err := strconv.ParseFloat(c.PostForm("Jumlah"), 64)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		aset_tik.Jumlah = int(jumlahFloat)
// 		satuan := c.PostForm("Satuan")
// 		aset_tik.Satuan = satuan
// 		keterangan := c.PostForm("Keterangan")
// 		aset_tik.Keterangan = keterangan
// 		aset_tik.Path = dbPath
// 		aset_tik.Gambar = fileName
// 		aset_tik.CreatedAt = time.Now()
// 		aset_tik.UpdatedAt = time.Now()

// 		success := asettikmodel.Create(&aset_tik)
// 		if success == nil {
// 			c.Redirect(http.StatusSeeOther, "/aset/aset-tik")
// 		} else {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create aset_tik"})
// 		}
// 	}
// }

// func Edit(c *gin.Context) {
// 	idString := c.Query("id")
// 	if idString == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing id parameter"})
// 		return
// 	}

// 	categories, _ := categorymodel.GetAll()
// 	tipes, _ := tipemodel.GetAll()
// 	aset_tik, err := asettikmodel.Detail(idString)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	path := map[string]string{
// 		"menu":    "aset",
// 		"subMenu": "aset-tik",
// 	}

// 	data := map[string]interface{}{
// 		"Title":        "Edit Aset TIK",
// 		"path":         path,
// 		"aset_tik":     aset_tik,
// 		"categories":   categories,
// 		"tipes":        tipes,
// 		"SelectedTipe": aset_tik.TipeID,
// 		"SelectedAset": aset_tik.KategoriID, // untuk selected kategori aset
// 	}

// 	c.HTML(http.StatusOK, "/aset/aset_tik/edit.html", data)
// }

// func Update(c *gin.Context) {
// 	idString := c.PostForm("ID")
// 	if idString == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing ID parameter"})
// 		return
// 	}

// 	aset, err := asettikmodel.Detail(idString)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	oldFilePath := filepath.Join(".", aset.Path)
// 	var aset_tik entities.AsetTik
// 	var dbPath, fileName string

// 	// check if a new file is uploaded
// 	file, _, err := c.Request.FormFile("Gambar")
// 	if err != nil {
// 		if err == http.ErrMissingFile {
// 			// Tidak ada file baru, lanjutkan tanpa mengubah file lama
// 			file = nil
// 			dbPath = aset.Path
// 			fileName = aset.Gambar
// 		} else {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving file"})
// 			return
// 		}
// 	} else {
// 		// Hapus file lama jika ada
// 		if _, err := os.Stat(oldFilePath); err == nil {
// 			err := os.Remove(oldFilePath)
// 			if err != nil {
// 				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete old file"})
// 				return
// 			}
// 		}

// 		// Simpan file baru
// 		path := "./public/uploads/aset/tetap/"
// 		fileName = helpers.RandString(30) + ".jpg"
// 		newFilePath := filepath.Join(path, fileName)
// 		dbPath = strings.ReplaceAll(filepath.Join("/public/uploads/aset/tetap/", fileName), "\\", "/")

// 		out, err := os.Create(newFilePath)

// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create new file"})
// 			return
// 		}
// 		defer out.Close()

// 		_, err = io.Copy(out, file)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save new file"})
// 			return
// 		}
// 		defer file.Close()
// 	}

// 	aset_tik.NamaAset = c.PostForm("NamaAset")
// 	aset_tik.JenisAset = c.PostForm("JenisAset")
// 	aset_tik.Merek = c.PostForm("Merek")
// 	aset_tik.Model = c.PostForm("Model")
// 	aset_tik.TanggalPerolehan, _ = time.Parse("2006-01-02", c.PostForm("TanggalPerolehan"))
// 	aset_tik.Nilai, _ = helpers.ParseCurrencyToFloat(c.PostForm("Nilai"))
// 	deskripsi := c.PostForm("Deskripsi")
// 	aset_tik.Deskripsi = deskripsi
// 	aset_tik.Path = dbPath
// 	aset_tik.Gambar = fileName
// 	aset_tik.KodeAset = c.PostForm("KodeAset")
// 	aset_tik.KategoriID = c.PostForm("KategoriID")
// 	aset_tik.TipeID = c.PostForm("TipeID")
// 	aset_tik.SerialNumber = c.PostForm("SerialNumber")
// 	aset_tik.Status = c.PostForm("Status")
// 	jumlahFloat, err := strconv.ParseFloat(c.PostForm("Jumlah"), 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	aset_tik.Jumlah = int(jumlahFloat)
// 	satuan := c.PostForm("Satuan")
// 	aset_tik.Satuan = satuan
// 	keterangan := c.PostForm("Keterangan")
// 	aset_tik.Keterangan = keterangan
// 	aset_tik.CreatedAt = time.Now()
// 	aset_tik.UpdatedAt = time.Now()

// 	success := asettikmodel.Update(idString, aset_tik)
// 	if success == nil {
// 		c.Redirect(http.StatusSeeOther, "/aset/aset-tik")
// 	} else {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update aset_tik"})
// 	}
// }

// func Distribusi(c *gin.Context) {
// 	idString := c.Query("aset_id")
// 	if idString == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing id parameter"})
// 		return
// 	}

// 	asets, err := asettikmodel.GetAll()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	bidang, err := bidangmodel.GetAllBidang()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	pegawai, err := pegawaimodel.GetAll()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	aset, err := asettikmodel.Detail(idString)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	path := map[string]string{
// 		"menu":    "aset",
// 		"subMenu": "aset-tik",
// 	}

// 	data := gin.H{
// 		"Title":        "Distribusi Aset",
// 		"path":         path,
// 		"aset":         aset,
// 		"asets":        asets,
// 		"bidang":       bidang,
// 		"pegawai":      pegawai,
// 		"SelectedAset": aset.ID,
// 	}

// 	c.HTML(http.StatusOK, "aset/aset_tik/distribusi.html", data)
// }

// func Delete(c *gin.Context) {
// 	idString := c.Query("id")

// 	if idString == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing id parameter"})
// 		return
// 	}

// 	if err := asettikmodel.Delete(idString); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete aset_tik"})
// 		return
// 	}

// 	c.Redirect(http.StatusSeeOther, "/aset/aset-tik")
// }
