package riwayatasetcontroller

import (
	"fmt"
	"inventaris/entities"
	"inventaris/helpers/helpers"
	"inventaris/helpers/queryhelpers"
	"io/ioutil"
	"net/http"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {
	dir := "./data"

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		http.Error(w, "Failed to read directory :"+err.Error(), http.StatusInternalServerError)
		return
	}

	var fileNames []string
    for _, file := range files {
        if strings.HasSuffix(file.Name(), ".json") {
            nameWithoutExt := strings.TrimSuffix(file.Name(), ".json")
            fileNames = append(fileNames, nameWithoutExt)
        }
    }

	var asetTikList []entities.AsetTik
	for _, kode := range fileNames {
		asetTikList = append(asetTikList, entities.AsetTik{Kode_Aset: kode})
	}

	data_aset := queryhelpers.GetAset(asetTikList)

	fmt.Println(data_aset)

	path := map[string]string{
		"menu": "riwayat-aset",
	}

	data := map[string]any{
		"Title": "Riwayat Aset",
		"path":  path,
		"files": fileNames,
	}

	helpers.RenderTemplate(w, "riwayataset/index.html", data)
}
