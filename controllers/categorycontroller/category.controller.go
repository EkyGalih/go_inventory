package categorycontroller

import (
	"inventaris/entities"
	"inventaris/helpers/helpers"
	"inventaris/models/categorymodel"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	// tambahkan flashes alert
	session := sessions.Default(c)
	flashes := session.Flashes()
	session.Save()

	categories, _ := categorymodel.GetAllCategory()
	path := map[string]string{
		"menu":    "addons",
		"subMenu": "category",
	}

	data := map[string]any{
		"Title":      "Category",
		"path":       path,
		"categories": categories,
		"Flashes":    flashes,
	}

	helpers.RenderTemplate(c, "addons/category/index.html", data)
}

func Add(c *gin.Context) {
	desc := c.PostForm("Deskripsi")
	category := entities.Category{
		NamaKategori: c.PostForm("NamaKategori"),
		Deskripsi:    &desc,
	}

	if err := categorymodel.CreateCategory(category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(c)
	session.AddFlash("Data berhasil disimpan")
	session.Save()

	c.Redirect(http.StatusSeeOther, "/addons/kategori")
}

func Edit(c *gin.Context) {
	session := sessions.Default(c)

	if c.Request.Method == http.MethodGet {
		id := c.Param("id")
		if id == "" {
			session.AddFlash("Id Kosong")
			session.Save()
		}
		category, err := categorymodel.GetCategoryByID(id)
		if err != nil {
			session.AddFlash("Category tidak ditemukan")
			session.Save()
		}

		data := map[string]interface{}{
			"Title": "Edit Category",
			"path": map[string]string{
				"menu":    "addons",
				"subMenu": "category",
			},
			"category": category,
		}

		helpers.RenderTemplate(c, "/addons/category/edit.html", data)
	}

	if c.Request.Method == http.MethodPost {
		idString := c.Param("id")
		if idString == "" {
			session = sessions.Default(c)
			session.AddFlash("Id Kosong")
			session.Save()
			c.Redirect(http.StatusFound, "/addons/category/edit.html")
			return
		}

		var category entities.Category
		if err := c.ShouldBind(&category); err != nil {
			session = sessions.Default(c)
			session.AddFlash("Gagal")
			session.Save()
			c.Redirect(http.StatusFound, "/addons/category/edit.html")
			return
		}

		category.ID = idString
		category.UpdatedAt = time.Now()

		err := categorymodel.UpdateCategory(category)
		if err != nil {
			session = sessions.Default(c)
			session.AddFlash("Error :", err.Error())
			session.Save()
			c.Redirect(http.StatusFound, "/addons/category/edit.html")
			return
		}

		session = sessions.Default(c)
		session.AddFlash("Data berhasil disimpan")
		session.Save()
		c.Redirect(http.StatusSeeOther, "/addons/kategori")
	}
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	categorymodel.DeleteCategory(id)
	session := sessions.Default(c)
	session.AddFlash("Data berhasil dihapus")
	session.Save()
	c.Redirect(http.StatusFound, "/addons/kategori")
}
