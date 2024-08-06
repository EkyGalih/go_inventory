package asettikmodel

import (
	"database/sql"
	"fmt"
	"inventaris/config"
	"inventaris/entities"

	"github.com/google/uuid"
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

func Create(aset_tik entities.AsetTik) (bool, error) {
	newUUID := uuid.New()
	result, err := config.DB.Exec(`
	INSERT INTO aset_tik (id, kode_aset, nama_aset, merek, model, serial_number, deskripsi, kategori_id, tanggal_perolehan, status, nilai, jumlah, keterangan, gambar, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		newUUID,
		aset_tik.Kode_Aset,
		aset_tik.Nama_Aset,
		aset_tik.Merek,
		aset_tik.Model,
		aset_tik.Serial_Number,
		aset_tik.Deskripsi,
		aset_tik.Kategori_id,
		aset_tik.Tanggal_Perolehan,
		aset_tik.Status,
		aset_tik.Nilai,
		aset_tik.Jumlah,
		aset_tik.Keterangan,
		aset_tik.Gambar,
		aset_tik.Created_At,
		aset_tik.Updated_At,
	)

	if err != nil {
		panic(err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	return rowsAffected > 0, nil
}

func Detail(id string) (entities.AsetTik, error) {
	row := config.DB.QueryRow(`Select * from aset_tik where id = ?`, id)

	var aset_tik entities.AsetTik
	if err := row.Scan(&aset_tik.Id, &aset_tik.Kode_Aset, &aset_tik.Nama_Aset, &aset_tik.Merek, &aset_tik.Model, &aset_tik.Serial_Number, &aset_tik.Deskripsi, &aset_tik.Kategori_id, &aset_tik.Tanggal_Perolehan, &aset_tik.Status, &aset_tik.Nilai, &aset_tik.Jumlah, &aset_tik.Keterangan, &aset_tik.Gambar, &aset_tik.Created_At, &aset_tik.Updated_At); err != nil {
		if err == sql.ErrNoRows {
			return aset_tik, fmt.Errorf("no category found with id %s", id)
		}
		return aset_tik, fmt.Errorf("failed to retrieve category: %w", err)
	}

	return aset_tik, nil
}
