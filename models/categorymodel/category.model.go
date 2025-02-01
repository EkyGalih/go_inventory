package categorymodel

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

const filePath = "data/kategori/kategori.json"

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

func loadFromFile() ([]entities.Category, error) {
	if err := ensureFileExists(); err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var categories []entities.Category
	if err := json.Unmarshal(data, &categories); err != nil {
		return nil, err
	}
	return categories, nil
}

func saveToFile(categories []entities.Category) error {
	if err := ensureFileExists(); err != nil {
		return err
	}
	data, err := json.MarshalIndent(categories, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, 0644)
}

func GetAllCategory() ([]entities.Category, error) {
	return loadFromFile()
}

func GetCategoryByID(id string) (*entities.Category, error) {
	categories, err := loadFromFile()
	if err != nil {
		return nil, err
	}

	for _, b := range categories {
		if b.ID == id {
			return &b, nil
		}
	}

	return nil, errors.New("Kategori tidak ditemukan")
}

func CreateCategory(category entities.Category) error {
	categories, err := loadFromFile()
	if err != nil {
		return err
	}

	newID := uuid.NewString()
	category.ID = newID
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()

	categories = append(categories, category)

	return saveToFile(categories)
}

func UpdateCategory(updatedCategory entities.Category) error {
	categories, err := loadFromFile()
	if err != nil {
		return err
	}

	found := false
	for i, category := range categories {
		if category.ID == updatedCategory.ID {
			updatedCategory.UpdatedAt = time.Now()
			categories[i] = updatedCategory
			found = true
			break
		}
	}

	if !found {
		return errors.New("category not found")
	}

	return saveToFile(categories)
}

func DeleteCategory(id string) error {
	categories, err := loadFromFile()
	if err != nil {
		return err
	}

	var deleteCategory []entities.Category
	found := false
	for _, b := range categories {
		if b.ID != id {
			deleteCategory = append(deleteCategory, b)
		} else {
			found = true
		}
	}

	if !found {
		return errors.New("Category tidak ditemukan")
	}

	return saveToFile(deleteCategory)
}
