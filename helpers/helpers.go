// helpers/helpers.go
package helpers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
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

func ParseCurrencyToFloat(value string) (float64, error) {
    // Menggunakan ekspresi reguler untuk menghapus karakter non-numerik
    re := regexp.MustCompile(`[^\d]`)
    cleaned := re.ReplaceAllString(value, "")

    // Mengonversi string yang sudah dibersihkan menjadi float
    number, err := strconv.ParseFloat(cleaned, 64)
    if err != nil {
        return 0, err
    }
    return number, nil
}

func FormatCurrency(value float64) string {
	// Convert the number to a string with no decimal places
	valueStr := strconv.FormatFloat(value, 'f', 0, 64)
	// Add periods as thousands separators
	var result strings.Builder
	n := len(valueStr)
	for i, c := range valueStr {
		if (n-i)%3 == 0 && i > 0 {
			result.WriteRune('.')
		}
		result.WriteRune(c)
	}
	return "Rp. " + result.String()
}