// helpers/helpers.go
package helpers

import (
    "html/template"
    "net/http"
    "path/filepath"
)

// RenderTemplate renders a template with the given name and data
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    t := template.New("").Funcs(template.FuncMap{})

    // Parse the partial templates
    t, err := t.ParseGlob("views/partials/*.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Parse the main template
    tmplPath := filepath.Join("views", tmpl)
    t, err = t.ParseFiles(tmplPath)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = t.ExecuteTemplate(w, filepath.Base(tmplPath), data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
