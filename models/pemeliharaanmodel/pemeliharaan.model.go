package pemeliharaanmodel

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"inventaris/entities"

	"github.com/google/uuid"
)

const filePath = "data/pemeliharan/pemeliharan.json"

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

func loadFromFile() ([]entities.Pemeliharaan, error) {
	if err := ensureFileExists(); err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var pemeliharan []entities.Pemeliharaan
	if err := json.Unmarshal(data, &pemeliharan); err != nil {
		return nil, err
	}
	return pemeliharan, nil
}

func saveToFile(pemeliharan []entities.Pemeliharaan) error {
	if err := ensureFileExists(); err != nil {
		return err
	}
	data, err := json.MarshalIndent(pemeliharan, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, 0644)
}

func GetAllPemeliharaan() ([]entities.Pemeliharaan, error) {
	return loadFromFile()
}

func GetPemeliharaanByID(id string) (*entities.Pemeliharaan, error) {
	pemeliharan, err := loadFromFile()
	if err != nil {
		return nil, err
	}

	for _, p := range pemeliharan {
		if p.ID == id {
			return &p, nil
		}
	}

	return nil, errors.New("Pemeliharaan tidak ditemukan")
}

func CreatePemeliharaan(pemeliharan entities.Pemeliharaan) error {
	pemeliharans, err := loadFromFile()
	if err != nil {
		return err
	}

	newID := uuid.NewString()
	pemeliharan.ID = newID
	pemeliharan.CreatedAt = time.Now()
	pemeliharan.UpdatedAt = time.Now()

	pemeliharans = append(pemeliharans, pemeliharan)

	return saveToFile(pemeliharans)
}

func UpdatePemeliharaan(updatedPemeliharaan entities.Pemeliharaan) error {
	pemeliharans, err := loadFromFile()
	if err != nil {
		return err
	}

	found := false
	for i, Pemeliharaan := range pemeliharans {
		if Pemeliharaan.ID	 == updatedPemeliharaan.ID {
			updatedPemeliharaan.UpdatedAt = time.Now()
			pemeliharans[i] = updatedPemeliharaan
			found = true
			break
		}
	}

	if !found {
		return errors.New("Pemeliharaan tidak ditemukan")
	}

	return saveToFile(pemeliharans)
}

func DeletePemeliharaan(id string) error {
	pemeliharans, err := loadFromFile()
	if err != nil {
		return err
	}

	var DeletePemeliharaan []entities.Pemeliharaan
	found := false
	for _, b := range pemeliharans {
		if b.ID != id {
			DeletePemeliharaan = append(DeletePemeliharaan, b)
		} else {
			found = true
		}
	}

	if !found {
		return errors.New("Pemeliharaan tidak ditemukan")
	}

	return saveToFile(DeletePemeliharaan)
}
