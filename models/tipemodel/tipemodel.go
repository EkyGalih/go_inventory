package tipemodel

import (
	"fmt"
	"inventaris/config"
	"inventaris/entities"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

func GetAll() ([]entities.Tipe, error) {
	var tipes []entities.Tipe
	err := config.DB.Order("updated_at DESC").Find(&tipes).Error
	if err != nil {
		return nil, err
	}
	return tipes, nil
}

func Create(tipe entities.Tipe) (bool, error) {
	newUUID := uuid.New()
	tipe.Id = newUUID.String()
	err := config.DB.Create(&tipe).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func Detail(id string) (entities.Tipe, error) {
	var tipe entities.Tipe
	err := config.DB.First(&tipe, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return tipe, fmt.Errorf("tipe not found with id %s", id)
		}
		return tipe, fmt.Errorf("failed to retrieve tipe: %w", err)
	}
	return tipe, nil
}

func Update(id string, tipe entities.Tipe) (bool, error) {
	err := config.DB.Model(&tipe).Where("id = ?", id).Updates(tipe).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func Delete(id string) error {
	err := config.DB.Delete(&entities.Tipe{}, "id = ?", id).Error
	return err
}
