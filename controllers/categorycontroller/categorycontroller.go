package categorycontroller

import (
	"inventaris/entities"
	"inventaris/helpers"
	"inventaris/models/categorymodel"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categorymodel.GetAll()
	data := map[string]any{
		"Title":      "Category",
		"categories": categories,
	}

	helpers.RenderTemplate(w, "category/index.html", data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data := map[string]any{
			"Title": "Add Category",
		}
		helpers.RenderTemplate(w, "category/create.html", data)
	}

	if r.Method == http.MethodPost {
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
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idString := r.URL.Query().Get("id")
		if idString == "" {
			http.Error(w, "Missing id parameter", http.StatusBadRequest)
			return
		}

		category, err := categorymodel.Detail(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data := map[string]interface{}{
			"Title":    "Edit Category",
			"category": category,
		}

		helpers.RenderTemplate(w, "category/edit.html", data)
		return // Ensure the function exits after handling GET request
	}

	if r.Method == http.MethodPost {
		idString := r.FormValue("id")
		if idString == "" {
			http.Error(w, "Missing id parameter", http.StatusBadRequest)
			return
		}

		var category entities.Category
		category.Nama_Kategori = r.FormValue("nama_kategori")
		category.Updated_At = time.Now()

		if ok := categorymodel.Update(idString, category); !ok {
			http.Error(w, "Failed to update category", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	if idString == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	if err := categorymodel.Delete(idString); err != nil {
		http.Error(w, "Failed to delete category", http.StatusInternalServerError) // Use http.Error instead of panic
		return
	}

	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}
