package asetmodel

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"inventaris/entities"

	"github.com/google/uuid"
)

const filePath = "data/aset/aset.json"

func ensureFileExists() error {
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = file.WriteString("[]")
		return err
	}
	return nil
}

func loadFromFile() ([]entities.Aset, error) {
	if err := ensureFileExists(); err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var Aset []entities.Aset
	if err := json.Unmarshal(data, &Aset); err != nil {
		return nil, err
	}
	return Aset, nil
}

// Fungsi untuk membaca data kategori dari file JSON
func loadKategoriFromFile() ([]entities.Category, error) {
	file, err := os.ReadFile("data/kategori/kategori.json")
	if err != nil {
		return nil, err
	}

	var kategoriList []entities.Category
	err = json.Unmarshal(file, &kategoriList)
	if err != nil {
		return nil, err
	}

	return kategoriList, nil
}

// Fungsi untuk membaca data tipe dari file JSON
func loadTipeFromFile() ([]entities.Tipe, error) {
	file, err := os.ReadFile("data/tipe/tipe.json")
	if err != nil {
		return nil, err
	}

	var tipeList []entities.Tipe
	err = json.Unmarshal(file, &tipeList)
	if err != nil {
		return nil, err
	}

	return tipeList, nil
}

func saveToFile(Aset []entities.Aset) error {
	if err := ensureFileExists(); err != nil {
		return err
	}
	data, err := json.MarshalIndent(Aset, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, 0644)
}

func GetAllAset() ([]entities.Aset, error) {
	return loadFromFile()
}

func GetAsetByID(id string) (*entities.Aset, error) {
	AsetList, err := loadFromFile()
	if err != nil {
		return nil, err
	}

	KategoriList, _ := loadKategoriFromFile()
	TipeList, _ := loadTipeFromFile()

	for _, aset := range AsetList {
		if aset.ID == id {
			// Cari Kategori berdasarkan KategoriID
			for _, kategori := range KategoriList {
				if kategori.ID == aset.KategoriID {
					aset.Kategori = &kategori
					break
				}
			}

			// Cari Tipe berdasarkan TipeID
			for _, tipe := range TipeList {
				if tipe.ID == aset.TipeID {
					aset.Tipe = &tipe
					break
				}
			}

			// Debug apakah kategori masih nil
			if aset.Kategori == nil {
				log.Println("Kategori tidak ditemukan untuk aset ID:", aset.ID)
			}

			if aset.Tipe == nil {
				log.Println("Tipe tidak ditemukan untuk aset ID:", aset.ID)
			}

			return &aset, nil
		}
	}

	return nil, errors.New("Aset tidak ditemukan")
}

func CreateAset(asets entities.Aset) error {
	Aset, err := loadFromFile()
	if err != nil {
		return err
	}

	newID := uuid.NewString()
	asets.ID = newID
	asets.CreatedAt = time.Now()
	asets.UpdatedAt = time.Now()

	Aset = append(Aset, asets)

	return saveToFile(Aset)
}

func UpdateAset(updatedAset entities.Aset) error {
	asets, err := loadFromFile()
	if err != nil {
		return err
	}

	found := false
	for i, Aset := range asets {
		if Aset.ID == updatedAset.ID {
			updatedAset.UpdatedAt = time.Now()
			asets[i] = updatedAset
			found = true
			break
		}
	}

	if !found {
		return errors.New("Aset tidak ditemukan")
	}

	return saveToFile(asets)
}

func DeleteAset(id string) error {
	asets, err := loadFromFile()
	if err != nil {
		return err
	}

	var deleteAset []entities.Aset
	found := false
	for _, b := range asets {
		if b.ID != id {
			deleteAset = append(deleteAset, b)
		} else {
			found = true
		}
	}

	if !found {
		return errors.New("Aset tidak ditemukan")
	}

	return saveToFile(deleteAset)
}
