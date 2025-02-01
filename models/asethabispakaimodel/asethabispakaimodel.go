package asethabispakaimodel

import (
	"errors"
	"inventaris/config"
	"inventaris/entities"
	"os"

	"github.com/google/uuid"
)

func GetAll(page, limit int) ([]entities.AsetTik, error) {
	var asetTiks []entities.AsetTik
	offset := (page - 1) * limit

	err := config.DB.Where("jenis_aset = ?", "Habis Pakai").Order("updated_at DESC").Limit(limit).Offset(offset).Find(&asetTiks).Error
	if err != nil {
		panic(err)
	}

	return asetTiks, nil
}

func GetTotalRows() (int64, error) {
	var totalRows int64
	err := config.DB.Model(&entities.AsetTik{}).
	Where("jenis_aset = ?", "Hais Pakai").
	Count(&totalRows).Error

	return totalRows, err
}

func Create(asetTik *entities.AsetTik) error {
	asetTik.ID = uuid.New().String() // generate UUID
	return config.DB.Create(asetTik).Error // GORM memerlukan pointer untuk membuat entri baru
}

func Detail(id string) (entities.AsetTik, error) {
	var asetTik entities.AsetTik
	err := config.DB.Where("jenis_aset = ? AND id = ?", "Habis Pakai", id).First(&asetTik).Error
	if err != nil {
		return asetTik, errors.New("aset tidak ditemukan")
	}
	return asetTik, nil
}

func Update(id string, updateData entities.AsetTik) error {
	var asetTik entities.AsetTik
	result := config.DB.First(&asetTik, "id = ?", id)
	if result.Error != nil {
		return errors.New("aset tidak ditemukan")
	}

	// Update semua field kecuali ID, CreatedAt, dan DeletedAt
	return config.DB.Model(&asetTik).Updates(updateData).Error
}

func Delete(id string) error {
	var asetTik entities.AsetTik

	// Cari data terlebih dahulu
	result := config.DB.First(&asetTik, "id = ?", id)
	if result.Error != nil {
		return errors.New("aset tidak ditemukan")
	}

	// Jika ada path gambar, hapus file terkait
	if asetTik.Path != "" {
		if err := removeFile(asetTik.Path); err != nil {
			return errors.New("gagal menghapus gambar: " + err.Error())
		}
	}

	// Hapus secara soft delete
	return config.DB.Delete(&asetTik).Error
}

func removeFile(path string) error {
	if path != "" {
		err := os.Remove(path)
		if err != nil && !os.IsNotExist(err) {
			return err
		}
	}
	return nil
}