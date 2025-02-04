package pegawaimodel

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

const filePath = "data/pegawai/pegawai.json"

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

func loadFromFile() ([]entities.Pegawai, error) {
	if err := ensureFileExists(); err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var pegawai []entities.Pegawai
	if err := json.Unmarshal(data, &pegawai); err != nil {
		return nil, err
	}
	return pegawai, nil
}

func saveToFile(pegawai []entities.Pegawai) error {
	if err := ensureFileExists(); err != nil {
		return err
	}
	data, err := json.MarshalIndent(pegawai, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, 0644)
}

func GetAllPegawai() ([]entities.Pegawai, error) {
	return loadFromFile()
}

func GetPegawaiByID(id string) (*entities.Pegawai, error) {
	pegawai, err := loadFromFile()
	if err != nil {
		return nil, err
	}

	for _, b := range pegawai {
		if b.Id == id {
			return &b, nil
		}
	}

	return nil, errors.New("Pegawai tidak ditemukan")
}

func CreatePegawai(pegawai entities.Pegawai) error {
	pegawais, err := loadFromFile()
	if err != nil {
		return err
	}

	newID := uuid.NewString()
	pegawai.Id = newID
	pegawai.CreatedAt = time.Now()
	pegawai.UpdatedAt = time.Now()

	pegawais = append(pegawais, pegawai)

	return saveToFile(pegawais)
}

func UpdatePegawai(updatedPegawai entities.Pegawai) error {
	pegawais, err := loadFromFile()
	if err != nil {
		return err
	}

	found := false
	for i, Pegawai := range pegawais {
		if Pegawai.Id == updatedPegawai.Id {
			updatedPegawai.UpdatedAt = time.Now()
			pegawais[i] = updatedPegawai
			found = true
			break
		}
	}

	if !found {
		return errors.New("Pegawai tidak ditemukan")
	}

	return saveToFile(pegawais)
}

func DeletePegawai(id string) error {
	pegawais, err := loadFromFile()
	if err != nil {
		return err
	}

	var deletePegawai []entities.Pegawai
	found := false
	for _, b := range pegawais {
		if b.Id != id {
			deletePegawai = append(deletePegawai, b)
		} else {
			found = true
		}
	}

	if !found {
		return errors.New("Pegawai tidak ditemukan")
	}

	return saveToFile(deletePegawai)
}
