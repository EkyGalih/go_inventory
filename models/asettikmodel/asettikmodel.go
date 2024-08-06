package asettikmodel

import (
	"inventaris/config"
	"inventaris/entities"
)

func GetAll() []entities.AsetTik {
	rows, err := config.DB.Query(`SELECT * FROM aset_tik ORDER BY updated_at DESC`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var aset_tiks []entities.AsetTik

	for rows.Next() {
		var aset_tik entities.AsetTik
		if err := rows.Scan(&aset_tik.Id, &aset_tik.Kode_Aset, &aset_tik.Nama_Aset, &aset_tik.Merek, &aset_tik.Model, &aset_tik.Serial_Number, &aset_tik.Deskripsi, &aset_tik.Kategori_id, &aset_tik.Tanggal_Perolehan, &aset_tik.Status, &aset_tik.Nilai, &aset_tik.Jumlah, &aset_tik.Keterangan, &aset_tik.Gambar, &aset_tik.Created_At, &aset_tik.Updated_At); err != nil {
			panic(err)
		}

		aset_tiks = append(aset_tiks, aset_tik)
	}

	return aset_tiks
}
