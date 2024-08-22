package queryhelpers

import (
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

func GetAset(aset_tiks []entities.AsetTik) map[string]string {
	asets := make(map[string]string)
	for _, item := range aset_tiks {
		var data string
		err := config.DB.QueryRow("SELECT * FROM aset_tik WHERE kode_aset = ?", item.Kode_Aset).Scan(&data)
		if err != nil {
			continue
		}
		asets[item.Kode_Aset] = data
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
