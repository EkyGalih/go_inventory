package asettikmodel

import (
	"errors"
	"inventaris/config"
	"inventaris/entities"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GetAll - Ambil semua data aset tetap
func GetAll() ([]entities.AsetTik, error) {
	var asetTiks []entities.AsetTik
	err := config.DB.Where("jenis_aset = ?", "Tetap").Order("updated_at DESC").Find(&asetTiks).Error
	return asetTiks, err
}

// GetPaginate - Ambil data aset tetap dengan pagination
func GetPaginate(page, limit int) ([]entities.AsetTik, error) {
	var asetTiks []entities.AsetTik
	offset := (page - 1) * limit
	err := config.DB.Where("jenis_aset = ?", "Tetap").Order("updated_at DESC").Limit(limit).Offset(offset).Find(&asetTiks).Error
	return asetTiks, err
}

// GetTotalRows - Hitung jumlah total data
func GetTotalRows() (int64, error) {
	var totalRows int64
	err := config.DB.Model(&entities.AsetTik{}).Where("jenis_aset = ?", "Tetap").Count(&totalRows).Error
	return totalRows, err
}

// Create - Tambah aset baru
func Create(asetTik *entities.AsetTik) error {
	asetTik.ID = uuid.New().String()
	return config.DB.Create(asetTik).Error
}

// Detail - Ambil detail aset berdasarkan ID
func Detail(id string) (entities.AsetTik, error) {
	var asetTik entities.AsetTik
	err := config.DB.Where("jenis_aset = ? AND id = ?", "Tetap", id).First(&asetTik).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return asetTik, errors.New("aset tidak ditemukan")
		}
		return asetTik, err
	}
	return asetTik, nil
}

// Update - Update data aset berdasarkan ID
func Update(id string, updateData entities.AsetTik) error {
	var asetTik entities.AsetTik

	if err := config.DB.First(&asetTik, "id = ?", id).Error; err != nil {
		return errors.New("aset tidak ditemukan")
	}

	return config.DB.Model(&asetTik).Updates(updateData).Error
}

// Delete - Hapus aset berdasarkan ID
func Delete(id string) error {
	var asetTik entities.AsetTik

	if err := config.DB.First(&asetTik, "id = ?", id).Error; err != nil {
		return errors.New("aset tidak ditemukan")
	}

	// Hapus file gambar jika ada
	if asetTik.Path != "" {
		if err := os.Remove(filepath.Join("./", asetTik.Path)); err != nil && !os.IsNotExist(err) {
			return errors.New("gagal menghapus gambar: " + err.Error())
		}
	}

	return config.DB.Delete(&asetTik).Error
}
