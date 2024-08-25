package riwayatasetcontroller

import (
	"encoding/json"
	"inventaris/entities"
	"inventaris/helpers/helpers"
	"inventaris/helpers/queryhelpers"
	"inventaris/models/asettikmodel"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {
	dir := "./data/riwayataset"

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

	dataAset := queryhelpers.GetAset(asetTikList)

	path := map[string]string{
		"menu": "riwayat-aset",
	}

	data := map[string]any{
		"Title":    "Riwayat Aset",
		"path":     path,
		"dataAset": dataAset,
	}

	helpers.RenderTemplate(w, "riwayataset/index.html", data)
}

func Show(w http.ResponseWriter, r *http.Request) {
	kodeAset := r.URL.Query().Get("kode_aset")
	dir := "./data/riwayataset/"
	filePath := filepath.Join(dir, kodeAset+".json")

	jsonFile, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "Failed to open file : "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer jsonFile.Close()

	//membaca seluruh isi file json
	jsonEncode, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		http.Error(w, "Failed to read file :"+err.Error(), http.StatusInternalServerError)
		return
	}

	var riwayat []entities.Riwayat
	err = json.Unmarshal(jsonEncode, &riwayat)
	if err != nil {
		http.Error(w, "Failed to parse JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	sort.Slice(riwayat, func(i, j int) bool {
		return riwayat[i].Tanggal_Aksi.After(riwayat[j].Tanggal_Aksi)
	})

	aset, err := asettikmodel.GetAsetByKode(kodeAset)
	if err != nil {
		http.Error(w, "Failed to get aset: "+err.Error(), http.StatusInternalServerError)
		return
	}

	path := map[string]string{
		"menu": "riwayat-aset",
	}

	data := map[string]any{
		"Title":   "Riwayat Aset",
		"path":    path,
		"aset":    aset,
		"riwayat": riwayat,
	}

	helpers.RenderTemplate(w, "riwayataset/logs.html", data)
}
