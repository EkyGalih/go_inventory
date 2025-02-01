package lokasiasetmodel

import (
	"fmt"
	"inventaris/config"
	"inventaris/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetAll() ([]entities.LokasiAset, error) {
	var lokasiAsets []entities.LokasiAset
	err := config.DB.Preload("Aset").Preload("Bidang").Preload("Pegawai").Find(&lokasiAsets).Error
	if err != nil {
		return nil, err
	}
	return lokasiAsets, nil
}

func Create(lokasiaset entities.LokasiAset) (bool, error) {
	lokasiaset.ID = uuid.New().String()
	err := config.DB.Create(&lokasiaset).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func Detail(id string) (entities.LokasiAset, error) {
	var lokasiaset entities.LokasiAset
	err := config.DB.Preload("Aset").Preload("Bidang").Preload("Pegawai").First(&lokasiaset, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return lokasiaset, fmt.Errorf("lokasi aset tidak ditemukan dengan id %s", id)
		}
		return lokasiaset, fmt.Errorf("failed to retrieve lokasi aset: %w", err)
	}
	return lokasiaset, nil
}

func DaftarAset(pegawai_id string) ([]entities.LokasiAset, error) {
	var lokasiasets []entities.LokasiAset
	err := config.DB.Preload("Aset").Preload("Bidang").Preload("Pegawai").Where("pegawai_id = ?", pegawai_id).Order("updated_at desc").Find(&lokasiasets).Error
	if err != nil {
		return nil, err
	}
	return lokasiasets, nil
}

func Update(id string, lokasiaset entities.LokasiAset) (bool, error) {
	err := config.DB.Model(&lokasiaset).Where("id = ?", id).Updates(map[string]interface{}{
		"aset_id":           lokasiaset.AsetID,
		"bidang_id":         lokasiaset.BidangID,
		"pegawai_id":        lokasiaset.PegawaiID,
		"tanggal_perolehan": lokasiaset.TanggalPerolehan,
		"tanggal_selesai":   lokasiaset.TanggalSelesai,
		"jenis_pemanfaatan": lokasiaset.JenisPemanfaatan,
		"keterangan":        lokasiaset.Keterangan,
		"updated_at":        lokasiaset.UpdatedAt,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func Delete(id string) error {
	err := config.DB.Delete(&entities.LokasiAset{}, "id = ?", id).Error
	return err
}
