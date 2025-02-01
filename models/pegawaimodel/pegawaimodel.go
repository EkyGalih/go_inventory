package pegawaimodel

import (
	"fmt"
	"inventaris/config"
	"inventaris/entities"
	"gorm.io/gorm"
)

func GetAll() ([]entities.Pegawai, error) {
	var pegawais []entities.Pegawai
	err := config.DB.Find(&pegawais).Error
	if err != nil {
		return nil, err
	}
	return pegawais, nil
}

func Detail(id string) (entities.Pegawai, error) {
	var pegawai entities.Pegawai
	err := config.DB.First(&pegawai, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return pegawai, fmt.Errorf("pegawai not found with id %s", id)
		}
		return pegawai, fmt.Errorf("failed to retrieve pegawai: %w", err)
	}
	return pegawai, nil
}
