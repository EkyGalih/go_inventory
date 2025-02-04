package tipemodel

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

const filePath = "data/tipe/tipe.json"

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

func loadFromFile() ([]entities.Tipe, error) {
	if err := ensureFileExists(); err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var tipe []entities.Tipe
	if err := json.Unmarshal(data, &tipe); err != nil {
		return nil, err
	}
	return tipe, nil
}

func saveToFile(tipe []entities.Tipe) error {
	if err := ensureFileExists(); err != nil {
		return err
	}
	data, err := json.MarshalIndent(tipe, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, 0644)
}

func GetAllTipe() ([]entities.Tipe, error) {
	return loadFromFile()
}

func GetTipeByID(id string) (*entities.Tipe, error) {
	tipe, err := loadFromFile()
	if err != nil {
		return nil, err
	}

	for _, b := range tipe {
		if b.ID == id {
			return &b, nil
		}
	}

	return nil, errors.New("Tipe tidak ditemukan")
}

func CreateTipe(tipe entities.Tipe) error {
	tipes, err := loadFromFile()
	if err != nil {
		return err
	}

	newID := uuid.NewString()
	tipe.ID = newID
	tipe.CreatedAt = time.Now()
	tipe.UpdatedAt = time.Now()

	tipes = append(tipes, tipe)

	return saveToFile(tipes)
}

func UpdateTipe(updatedTipe entities.Tipe) error {
	categories, err := loadFromFile()
	if err != nil {
		return err
	}

	found := false
	for i, category := range categories {
		if category.ID == updatedTipe.ID {
			updatedTipe.UpdatedAt = time.Now()
			categories[i] = updatedTipe
			found = true
			break
		}
	}

	if !found {
		return errors.New("Tipe tidak ditemukan")
	}

	return saveToFile(categories)
}

func DeleteTipe(id string) error {
	tipes, err := loadFromFile()
	if err != nil {
		return err
	}

	var deleteTipe []entities.Tipe
	found := false
	for _, b := range tipes {
		if b.ID != id {
			deleteTipe = append(deleteTipe, b)
		} else {
			found = true
		}
	}

	if !found {
		return errors.New("Tipe tidak ditemukan")
	}

	return saveToFile(deleteTipe)
}
