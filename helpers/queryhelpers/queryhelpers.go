package queryhelpers

import (
	"database/sql"
	"fmt"
	"inventaris/config"
	"inventaris/entities"
)

// GetDistribusi generates a map of asset distribution based on the provided asset list.
//
// Parameter aset_tiks is a list of AsetTik entities.
// Returns a map of asset IDs to their respective distribution counts.
func GetDistribusi(aset_tiks []entities.AsetTik) map[string]int {
	distribusi := make(map[string]int)
	for _, aset := range aset_tiks {
		var count int
		err := config.DB.QueryRow("SELECT COUNT(*) FROM lokasi_aset WHERE aset_id = ?", aset.Id).Scan(&count)
		if err != nil {
			// Handle error, misalnya dengan logging
			continue
		}
		distribusi[aset.Id] = count
	}
	return distribusi
}

func GetAset(aset_tiks []entities.AsetTik) map[string]entities.AsetTik {
	asets := make(map[string]entities.AsetTik)
	for _, item := range aset_tiks {
		var aset entities.AsetTik
		query := "SELECT id, nama_aset, kode_aset, path FROM aset_tik WHERE kode_aset = ?"
		err := config.DB.QueryRow(query, item.Kode_Aset).Scan(&aset.Id, &aset.Nama_Aset, &aset.Kode_Aset, &aset.Path)

		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Printf("no aset found for kode aset: %s\n", item.Kode_Aset)
			} else {
				fmt.Printf("Failed to query aset for kode_aset: %s, error: %v\n", item.Kode_Aset, err)
			}
			continue
		}	
		asets[item.Kode_Aset] = aset
	}

	return asets
}

func CountAsetPegawai(lokasiaset []entities.LokasiAset) map[string]int {
	asetPegawais := make(map[string]int)
	for _, aset := range lokasiaset {
		var count int
		err := config.DB.QueryRow("SELECT COUNT(*) FROM lokasi_aset WHERE pegawai_id = ?", aset.Pegawai_id).Scan(&count)
		if err != nil {
			continue
		}
		asetPegawais[aset.Pegawai_id] = count
	}
	return asetPegawais
}
