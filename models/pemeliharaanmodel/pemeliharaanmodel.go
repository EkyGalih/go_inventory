package pemeliharaanmodel

import (
	"fmt"
	"inventaris/config"
	"inventaris/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetAll() ([]entities.Pemeliharaan, error) {
	var pemeliharaans []entities.Pemeliharaan
	err := config.DB.Table("pemeliharaan_aset").
		Select("pa.*, at.nama_aset, at.kode_aset, at.path").
		Joins("JOIN aset_tik at ON pa.aset_id = at.id").
		Order("pa.updated_at DESC").
		Find(&pemeliharaans).Error
	if err != nil {
		return nil, err
	}
	return pemeliharaans, nil
}

func Create(pemeliharaan entities.Pemeliharaan) (bool, error) {
	newUUID := uuid.New()
	pemeliharaan.Id = newUUID.String()
	err := config.DB.Create(&pemeliharaan).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func Detail(id string) (entities.Pemeliharaan, error) {
	var pemeliharaan entities.Pemeliharaan
	err := config.DB.Table("pemeliharaan_aset").
		Select("pa.*, at.nama_aset, at.kode_aset, at.path").
		Joins("JOIN aset_tik at ON pa.aset_id = at.id").
		Where("pa.id = ?", id).
		Order("pa.updated_at DESC").
		First(&pemeliharaan).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return pemeliharaan, fmt.Errorf("pemeliharaan aset tidak ditemukan dengan id %s", id)
		}
		return pemeliharaan, fmt.Errorf("failed to retrieve pemeliharaan aset: %w", err)
	}
	return pemeliharaan, nil
}

func Update(id string, pemeliharaan entities.Pemeliharaan) (bool, error) {
	err := config.DB.Model(&pemeliharaan).
		Where("id = ?", id).
		Updates(pemeliharaan).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func ChangeStatus(id string, pemeliharaan entities.Pemeliharaan) (bool, error) {
	err := config.DB.Model(&pemeliharaan).
		Where("id = ?", id).
		Update("status", pemeliharaan.Status).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func Delete(id string) error {
	err := config.DB.Delete(&entities.Pemeliharaan{}, "id = ?", id).Error
	return err
}
