package bidangmodel

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

const filePath = "data/bidang/bidang.json"

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

func loadFromFile() ([]entities.Bidang, error) {
	if err := ensureFileExists(); err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var bidangs []entities.Bidang
	if err := json.Unmarshal(data, &bidangs); err != nil {
		return nil, err
	}
	return bidangs, nil
}

func saveToFile(bidangs []entities.Bidang) error {
	if err := ensureFileExists(); err != nil {
		return err
	}
	data, err := json.MarshalIndent(bidangs, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, 0644)
}

func GetAllBidang() ([]entities.Bidang, error) {
	return loadFromFile()
}

func GetBidangByID(id string) (*entities.Bidang, error) {
	bidangs, err := loadFromFile()
	if err != nil {
		return nil, err
	}

	for _, b := range bidangs {
		if b.ID == id {
			return &b, nil
		}
	}

	return nil, errors.New("Bidang tidak ditemukan")
}

func CreateBidang(namaBidang string) error {
	bidangs, err := loadFromFile()
	if err != nil {
		return err
	}

	newID := uuid.NewString()
	bidangs = append(bidangs, entities.Bidang{ID: newID, NamaBidang: namaBidang, CreatedAt: time.Now(), UpdatedAt: time.Now()})

	return saveToFile(bidangs)
}

func UpdateBidang(id string, namaBidang string) error {
	bidangs, err := loadFromFile()
	if err != nil {
		return err
	}

	for i, b := range bidangs {
		if b.ID == id {
			bidangs[i].NamaBidang = namaBidang
			return saveToFile(bidangs)
		}
	}

	return errors.New("Bidang tidak ditemukan")
}

func DeleteBidang(id string) error {
	bidangs, err := loadFromFile()
	if err != nil {
		return err
	}

	var updatedBidangs []entities.Bidang
	found := false
	for _, b := range bidangs {
		if b.ID != id {
			updatedBidangs = append(updatedBidangs, b)
		} else {
			found = true
		}
	}

	if !found {
		return errors.New("Bidang tidak ditemukan")
	}

	return saveToFile(updatedBidangs)
}
