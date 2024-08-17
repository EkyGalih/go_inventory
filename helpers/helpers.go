// helpers/helpers.go
package helpers

import (
	"fmt"
	"html/template"
	"inventaris/config"
	"inventaris/entities"
	"math"
	"math/rand"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// RenderTemplate renders a Go template with the given data and writes the result to the HTTP response writer.
//
// The tmpl parameter specifies the path to the template file, and the data parameter is the data to be passed to the template.
// The function returns no value, but writes the rendered template to the HTTP response writer.
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    t := template.New("").Funcs(template.FuncMap{
        "mul": mul,
        "formatCurrency": FormatCurrency,
        "formatDate": formatDate,
        "floatToInt": ConvertFloatToInt,
        "removeHTMLTags": removeHTMLTags,
        "calculateAssetAge": calculateAssetAge,
        "add": add,
        "sub": sub,
        "until": until,
        "toInt": toInt,
        "StrLimit": StrLimit,
    })

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

// ParseCurrencyToFloat mengonversi string mata uang menjadi float64.
//
// Parameter value adalah string yang mewakili mata uang.
// Fungsi ini mengembalikan nilai float64 yang sudah dikonversi dan error jika terjadi kesalahan.
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

// FormatCurrency formats a float64 value as a currency string with periods as thousands separators.
//
// Parameters:
// - value: the float64 value to be formatted.
//
// Returns:
// - string: the formatted currency string.
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

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"


// RandString generates a random string of a specified length.
//
// Parameter n is the length of the string to be generated.
// Returns a string of random characters.
func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// mul returns the product of two float64 numbers.
//
// Parameters:
//  a (float64): the first number to be multiplied.
//  b (float64): the second number to be multiplied.
// Returns:
//  float64: the product of a and b.
func mul(a, b float64) float64 {
	return a * b
}

// ConvertFloatToInt converts a float64 number to an integer.
//
// Parameter f is the float64 number to be converted.
// Returns an integer representation of the input float64 number.
func ConvertFloatToInt(f float64) int {
    return int(math.Round(f))
}

// formatDate formats a given time.Time object into a string.
//
// Parameter t is the time.Time object to be formatted.
// Returns a string representing the formatted date in the format "YYYY-MM-DD".
func formatDate(t time.Time) string {
    return t.Format("2006-01-02")
}

// GetDistribusi generates a map of asset distribution based on the provided asset list.
//
// Parameter aset_tiks is a list of AsetTik entities.
// Returns a map of asset IDs to their respective distribution counts.
func GetDistribusi(aset_tiks []entities.AsetTik) map[string]int {
    distribusi := make(map[string]int)
    for _, aset := range aset_tiks {
        var count int
        err := config.DB.QueryRow("SELECT COUNT(*) FROM lokasi_aset WHERE aset_id = ?", aset.Id).Scan(&count)
        if err != nil {
            // Handle error, misalnya dengan logging
            continue
        }
        distribusi[aset.Id] = count
    }
    return distribusi
}

// removeHTMLTags removes HTML tags from a given input string.
//
// Parameter input is the string from which HTML tags will be removed.
// Returns the input string with all HTML tags removed.
func removeHTMLTags(input string) string {
	re := regexp.MustCompile("<.*?>")
	return re.ReplaceAllString(input, "")
}

func calculateAssetAge(tanggalPerolehan time.Time) string {
	today := time.Now()

	// Menghitung perbedaan tahun
	years := today.Year() - tanggalPerolehan.Year()

	// Menghitung perbedaan bulan
	months := int(today.Month()) - int(tanggalPerolehan.Month())
	if months < 0 {
		years--
		months += 12
	}

	// Membuat string hasilnya
	return fmt.Sprintf("%d tahun %d bulan", years, months)
}


// pagintation
func add(a, b int) int {
    return a+b
}

func sub(a, b int) int {
    return a-b
}

func until(count int) []int {
    var i int
    var items []int
    for i = 0; i < count; i++ {
        items = append(items, i)
    }
    return items
}

func toInt(val interface{}) int {
    switch v := val.(type) {
    case int:
        return v
    case string:
        if i, err := strconv.Atoi(v); err == nil {
            return i
        }
    }

    return 0
}

// end pagination

func StrLimit(s string, limit int) string {
	if len(s) > limit {
		return s[:limit] + "..."
	}
	return s
}