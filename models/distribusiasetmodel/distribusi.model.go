package distribusiasetmodel

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"inventaris/entities"

	"github.com/google/uuid"
)

const filePath = "data/distribusi/distribusi.json"

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

func loadFromFile() ([]entities.Distribusi, error) {
	if err := ensureFileExists(); err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var distribusi []entities.Distribusi
	if err := json.Unmarshal(data, &distribusi); err != nil {
		return nil, err
	}
	return distribusi, nil
}

func saveToFile(distribusi []entities.Distribusi) error {
	if err := ensureFileExists(); err != nil {
		return err
	}
	data, err := json.MarshalIndent(distribusi, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, 0644)
}

func GetAllDistribusi() ([]entities.Distribusi, error) {
	return loadFromFile()
}


func CreateDistribusi(dist entities.Distribusi) error {
	distribusi, err := loadFromFile()
	if err != nil {
		return err
	}

	newID := uuid.NewString()
	dist.ID = newID
	dist.CreatedAt = time.Now()
	dist.UpdatedAt = time.Now()

	distribusi = append(distribusi, dist)

	return saveToFile(distribusi)
}
