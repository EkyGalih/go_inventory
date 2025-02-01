package queryhelpers

import (
	"fmt"
	"inventaris/config"
	"inventaris/entities"

	"gorm.io/gorm"
)

// GetDistribusi generates a map of asset distribution based on the provided asset list.
//
// Parameter aset_tiks is a list of AsetTik entities.
// Returns a map of asset IDs to their respective distribution counts.
func GetDistribusi(aset_tiks []entities.AsetTik) map[string]int {
	distribusi := make(map[string]int)
	for _, aset := range aset_tiks {
		var count int64
		err := config.DB.Model(&entities.LokasiAset{}).Where("aset_id = ?", aset.ID).Count(&count).Error
		if err != nil {
			// Handle error, misalnya dengan logging
			continue
		}
		distribusi[aset.ID] = int(count)
	}
	return distribusi
}

func GetAset(aset_tiks []entities.AsetTik) map[string]entities.AsetTik {
	asets := make(map[string]entities.AsetTik)
	for _, item := range aset_tiks {
		var aset entities.AsetTik
		err := config.DB.Where("kode_aset = ?", item.KodeAset).First(&aset).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				fmt.Printf("no aset found for kode aset: %s\n", item.KodeAset)
			} else {
				fmt.Printf("Failed to query aset for kode_aset: %s, error: %v\n", item.KodeAset, err)
			}
			continue
		}
		asets[item.KodeAset] = aset
	}

	return asets
}

func CountAsetPegawai(lokasiaset []entities.LokasiAset) map[string]int {
	asetPegawais := make(map[string]int)
	for _, aset := range lokasiaset {
		var count int64
		err := config.DB.Model(&entities.LokasiAset{}).Where("pegawai_id = ?", aset.PegawaiID).Count(&count).Error
		if err != nil {
			continue
		}
		asetPegawais[aset.PegawaiID] = int(count)
	}
	return asetPegawais
}
