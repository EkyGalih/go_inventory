package pemeliharaanmodel

import (
	"database/sql"
	"fmt"
	"inventaris/config"
	"inventaris/entities"

	"github.com/google/uuid"
)

func GetAll() []entities.Pemeliharaan {
	rows, err := config.DB.Query(`
	SELECT 
		pa.*, 
		at.nama_aset, 
		at.kode_aset ,
		at.path
	FROM 
		pemeliharaan_aset pa
	JOIN 
		aset_tik at ON pa.aset_id = at.id
	ORDER BY 
		pa.updated_at DESC;

	`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var pemeliharaans []entities.Pemeliharaan

	for rows.Next() {
		var pemeliharaan entities.Pemeliharaan

		if err := rows.Scan(
			&pemeliharaan.Id,
			&pemeliharaan.Aset_id,
			&pemeliharaan.Tanggal_Pemeliharaan,
			&pemeliharaan.Kerusakan,
			&pemeliharaan.Perbaikan,
			&pemeliharaan.Keterangan,
			&pemeliharaan.Status,
			&pemeliharaan.Nota,
			&pemeliharaan.Biaya,
			&pemeliharaan.Created_At,
			&pemeliharaan.Updated_At,
			&pemeliharaan.Nama_Aset,
			&pemeliharaan.Kode_Aset,
			&pemeliharaan.Path,
			); err != nil {
			panic(err.Error())
		}

		pemeliharaans = append(pemeliharaans, pemeliharaan)
	}

	return pemeliharaans
}

func Create(pemeliharaan entities.Pemeliharaan) (bool, error) {
	newUUID := uuid.New()
	result, err := config.DB.Exec(`
	INSERT INTO pemeliharaan_aset (id, aset_id, tanggal_pemeliharaan, kerusakan, perbaikan, keterangan, status, nota, biaya, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		newUUID,
		pemeliharaan.Aset_id,
		pemeliharaan.Tanggal_Pemeliharaan,
		pemeliharaan.Kerusakan,
		pemeliharaan.Perbaikan,
		pemeliharaan.Keterangan,
		pemeliharaan.Status,
		pemeliharaan.Nota,
		pemeliharaan.Biaya,
		pemeliharaan.Created_At,
		pemeliharaan.Updated_At,
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

func Detail(id string) (entities.Pemeliharaan, error) {
	row := config.DB.QueryRow(`
	SELECT 
		pa.*, 
		at.nama_aset, 
		at.kode_aset,
		at.path
	FROM 
		pemeliharaan_aset pa
	JOIN 
		aset_tik at ON pa.aset_id = at.id
	WHERE 
		pa.id = ?
	ORDER BY 
		pa.updated_at DESC`, id)

	var pemeliharaan entities.Pemeliharaan

	if err := row.Scan(&pemeliharaan.Id, 
		&pemeliharaan.Aset_id, 
		&pemeliharaan.Tanggal_Pemeliharaan, 
		&pemeliharaan.Kerusakan, 
		&pemeliharaan.Perbaikan, 
		&pemeliharaan.Keterangan, 
		&pemeliharaan.Status, 
		&pemeliharaan.Nota, 
		&pemeliharaan.Biaya, 
		&pemeliharaan.Created_At, 
		&pemeliharaan.Updated_At,
		&pemeliharaan.Nama_Aset,
		&pemeliharaan.Kode_Aset,
		&pemeliharaan.Path,
		); err != nil {
		if err == sql.ErrNoRows {
			return pemeliharaan, fmt.Errorf("pemeliharaan aset tidak ditemukan dengan id %s", id)
		}
		return pemeliharaan, fmt.Errorf("failet to retrieve pemeliharaan aset: %w", err)
	}

	return pemeliharaan, nil
}

func Update(id string, pemeliharaan entities.Pemeliharaan) (bool, error) {
	query, err := config.DB.Exec(`
	UPDATE pemeliharaan_aset
		SET
			aset_id = ?,
			tanggal_pemeliharaan = ?,
			kerusakan = ?,
			perbaikan = ?,
			keterangan = ?,
			status = ?,
			nota = ?,
			biaya = ?,
			updated_at = ?
		WHERE id = ?
	`,
		pemeliharaan.Aset_id,
		pemeliharaan.Tanggal_Pemeliharaan,
		pemeliharaan.Kerusakan,
		pemeliharaan.Perbaikan,
		pemeliharaan.Keterangan,
		pemeliharaan.Status,
		pemeliharaan.Nota,
		pemeliharaan.Biaya,
		pemeliharaan.Updated_At,
		id,
	)
	if err != nil {
		panic(err.Error())
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0, nil
}

func ChangeStatus(id string, pemeliharaan entities.Pemeliharaan) (bool, error) {
	query, err := config.DB.Exec(`
	UPDATE pemeliharaan_aset
		SET
		status = ?,
		updated_at = ?
		WHERE id = ?
	`,
		pemeliharaan.Status,
		pemeliharaan.Updated_At,
		id,
	)
	if err != nil {
		panic(err.Error())
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0, nil
}

func Delete(id string) error {
	_, err := config.DB.Exec(`DELETE FROM pemeliharaan_aset WHERE id = ?`, id)
	return err
}
