package asettikmodel

import (
	"database/sql"
	"errors"
	"fmt"
	"inventaris/config"
	"inventaris/entities"
	"os"

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
		if err := rows.Scan(&aset_tik.Id, &aset_tik.Kode_Aset, &aset_tik.Nama_Aset, &aset_tik.Merek, &aset_tik.Model, &aset_tik.Serial_Number, &aset_tik.Deskripsi, &aset_tik.Kategori_id, &aset_tik.Tanggal_Perolehan, &aset_tik.Status, &aset_tik.Nilai, &aset_tik.Jumlah, &aset_tik.Keterangan, &aset_tik.Path, &aset_tik.Gambar, &aset_tik.Created_At, &aset_tik.Updated_At); err != nil {
			panic(err)
		}

		aset_tiks = append(aset_tiks, aset_tik)
	}

	return aset_tiks
}

func Create(aset_tik entities.AsetTik) (bool, error) {
	newUUID := uuid.New()
	result, err := config.DB.Exec(`
	INSERT INTO aset_tik (id, kode_aset, nama_aset, merek, model, serial_number, deskripsi, kategori_id, tanggal_perolehan, status, nilai, jumlah, keterangan, path, gambar, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
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
		aset_tik.Path,
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
	if err := row.Scan(&aset_tik.Id, &aset_tik.Kode_Aset, &aset_tik.Nama_Aset, &aset_tik.Merek, &aset_tik.Model, &aset_tik.Serial_Number, &aset_tik.Deskripsi, &aset_tik.Kategori_id, &aset_tik.Tanggal_Perolehan, &aset_tik.Status, &aset_tik.Nilai, &aset_tik.Jumlah, &aset_tik.Keterangan, &aset_tik.Path, &aset_tik.Gambar, &aset_tik.Created_At, &aset_tik.Updated_At); err != nil {
		if err == sql.ErrNoRows {
			return aset_tik, fmt.Errorf("no category found with id %s", id)
		}
		return aset_tik, fmt.Errorf("failed to retrieve category: %w", err)
	}

	return aset_tik, nil
}

func Update(id string, aset_tik entities.AsetTik) bool {
	query, err := config.DB.Exec(`
	UPDATE aset_tik 
		SET 
			kode_aset = ?, 
			nama_aset = ?, 
			merek = ?, 
			model = ?, 
			serial_number = ?, 
			deskripsi = ?, 
			kategori_id = ?, 
			tanggal_perolehan = ?, 
			status = ?, 
			nilai = ?, 
			jumlah = ?, 
			keterangan = ?, 
			path = ?, 
			gambar = ?, 
			updated_at = ? 
		WHERE id = ?
	`,
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
		aset_tik.Path,
		aset_tik.Gambar,
		aset_tik.Updated_At,
		aset_tik.Id,
	)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id string) error {
	var filePath string

	err := config.DB.QueryRow(`SELECT path FROM aset_tik WHERE id = ?`, id).Scan(&filePath)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("no record found with the given id")
		}
		return err
	}
	
	if filePath != "" {
		err = os.Remove(filePath)
		if err != nil && !os.IsNotExist(err) {
			return errors.New("failed to delete image " + err.Error())
		}
	}
	
	_, err = config.DB.Exec("DELETE FROM aset_tik WHERE id = ?", id)
	if err != nil {
		return err
	}
	
	return nil
}
