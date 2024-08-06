package categorycontroller

import (
	"inventaris/entities"
	"inventaris/helpers"
	"inventaris/models/categorymodel"
	"net/http"
	"strconv"
	"time"
)


func Index(w http.ResponseWriter, r *http.Request) {
	categories := categorymodel.GetAll()
	data := map[string]any{
		"Title": "Category",
		"categories": categories,
	}

	helpers.RenderTemplate(w, "category/index.html", data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		helpers.RenderTemplate(w, "category/create.html", nil)
	}

	if r.Method == "POST" {
		var category entities.Category

		category.Nama_Kategori = r.FormValue("nama_kategori")
		category.Created_At = time.Now()
		category.Updated_At = time.Now()

		success, err := categorymodel.Create(category)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if success {
			http.Redirect(w, r, "/categories", http.StatusSeeOther)
		} else {
			http.Error(w, "Failed to create category", http.StatusInternalServerError)
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Invalid ID Format", http.StatusBadRequest)
		}

		category := categorymodel.Detail(id)
		data := map[string]any{
			"category": category,
		}

		helpers.RenderTemplate(w, "category/edit.html", data)
	}

	if (r.Method == "POST") {
		var category entities.Category
		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err.Error())
		}

		category.Nama_Kategori = r.FormValue("nama_kategori")
		category.Updated_At = time.Now()

		if ok := categorymodel.Update(id, category); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err.Error())
	}

	if err := categorymodel.Delete(id); err != nil {
		panic(err.Error())
	}

	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}
