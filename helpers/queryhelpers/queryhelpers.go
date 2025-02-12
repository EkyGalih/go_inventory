package queryhelpers

import (
	"encoding/json"
	"inventaris/entities"
	"inventaris/models/bidangmodel"
	"inventaris/models/pegawaimodel"
	"os"
	"path/filepath"
)

// GetDistribusi generates a map of asset distribution based on the provided asset list.
//
// Parameter aset_tiks is a list of AsetTik entities.
// Returns a map of asset IDs to their respective distribution counts.
func GetDistribusi(aset []entities.Aset) map[string]int {
	distribusi := make(map[string]int)
	
	filePath := filepath.Join("data", "distribusi", "distribusi.json")
	file, err := os.Open(filePath)
	if err != nil {
		return distribusi
	}
	defer file.Close()

	var distribusiData []entities.Distribusi
	err = json.NewDecoder(file).Decode(&distribusiData)
	if err != nil {
		return distribusi
	}

	for _, d := range distribusiData {
		distribusi[d.AsetID]++
	}

	return distribusi
}

func GetPemegangAsets(asets []entities.Aset) map[string][]entities.Distribusi {
	pemegangAsets := make(map[string][]entities.Distribusi)

	filePath := filepath.Join("data", "distribusi", "distribusi.json")
	file, err := os.Open(filePath)
	if err != nil {
		return pemegangAsets
	}
	defer file.Close()

	var distribusiData []entities.Distribusi
	err = json.NewDecoder(file).Decode(&distribusiData)
	if err != nil {
		return pemegangAsets
	}

	// load data bidang dan pegawai
	bidangs, _ := bidangmodel.GetAllBidang()
	pegawais, _ := pegawaimodel.GetAllPegawai()

	bidangMap := make(map[string]entities.Bidang)
	pegawaiMap := make(map[string]entities.Pegawai)

	for _, b := range bidangs {
		bidangMap[b.ID] = b
	}

	for _, p := range pegawais {
		pegawaiMap[p.Id] = p
	}

	for i, d := range distribusiData {
		if bidang, exists := bidangMap[d.BidangID]; exists {
			distribusiData[i].Bidang = &bidang
		}

		if d.PegawaiID != nil {
			if pegawai, exists := pegawaiMap[*d.PegawaiID]; exists {
				distribusiData[i].Pegawai = &pegawai
			}
		}
	}

	asetMap := make(map[string]bool)
	for _, aset := range asets {
		asetMap[aset.ID] = true
	}

	for _, d := range distribusiData {
		if asetMap[d.AsetID] {
			pemegangAsets[d.AsetID] = append(pemegangAsets[d.AsetID], d)
		}
	}

	return pemegangAsets
}

// func GetAset(aset_tiks []entities.Aset) map[string]entities.Aset {
// 	asets := make(map[string]entities.Aset)
// 	for _, item := range aset_tiks {
// 		var aset entities.Aset
// 		err := config.DB.Where("kode_aset = ?", item.KodeAset).First(&aset).Error
// 		if err != nil {
// 			if err == gorm.ErrRecordNotFound {
// 				fmt.Printf("no aset found for kode aset: %s\n", item.KodeAset)
// 			} else {
// 				fmt.Printf("Failed to query aset for kode_aset: %s, error: %v\n", item.KodeAset, err)
// 			}
// 			continue
// 		}
// 		asets[item.KodeAset] = aset
// 	}

// 	return asets
// }

// func CountAsetPegawai(lokasiaset []entities.Distribusi) map[string]int {
// 	asetPegawais := make(map[string]int)
// 	for _, aset := range lokasiaset {
// 		var count int64
// 		err := config.DB.Model(&entities.Distribusi{}).Where("pegawai_id = ?", aset.PegawaiID).Count(&count).Error
// 		if err != nil {
// 			continue
// 		}
// 		asetPegawais[*aset.PegawaiID] = int(count)
// 	}
// 	return asetPegawais
// }
