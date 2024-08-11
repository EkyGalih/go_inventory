package tipemodel

import (
	"database/sql"
	"fmt"
	"inventaris/config"
	"inventaris/entities"

	"github.com/google/uuid"
)

func GetAll() []entities.Tipe {
	rows, err := config.DB.Query(`SELECT * FROM tipe_aset ORDER BY updated_at DESC`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var tipes []entities.Tipe

	for rows.Next() {
		var tipe entities.Tipe
		if err := rows.Scan(&tipe.Id, &tipe.Nama_Tipe, &tipe.Keterangan, &tipe.Created_At, &tipe.Updated_At); err != nil {
			panic(err)
		}

		tipes = append(tipes, tipe)
	}

	return tipes
}

func Create(tipe entities.Tipe) (bool, error) {
	newUUID := uuid.New()
	result, err := config.DB.Exec(`
	INSERT INTO tipe_aset (id, nama_tipe, keterangan, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?)`,
		newUUID,
		tipe.Nama_Tipe,
		tipe.Keterangan,
		tipe.Created_At,
		tipe.Updated_At,
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

func Detail(id string) (entities.Tipe, error) {
	row := config.DB.QueryRow(`SELECT * FROM tipe_aset WHERE id = ?`, id)

	var tipes entities.Tipe
	if err := row.Scan(&tipes.Id, &tipes.Nama_Tipe, &tipes.Keterangan, &tipes.Created_At, &tipes.Updated_At); err != nil {
		if err == sql.ErrNoRows {
			return tipes, fmt.Errorf("tipe not found")
		}
		return tipes, fmt.Errorf("failed to retrieve tipe: %w", err)
	}

	return tipes, nil
}

func Update(id string, tipe entities.Tipe) bool {
	query, err := config.DB.Exec(`UPDATE tipe_aset SET nama_tipe = ?, keterangan = ?, updated_at = ? WHERE id = ?`, tipe.Nama_Tipe, tipe.Keterangan, tipe.Updated_At, id)
	if err != nil {
		panic(err.Error())
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	return result > 0
}

func Delete(id string) error {
	_, err := config.DB.Exec(`DELETE FROM tipe_aset WHERE id = ?`, id)
	return err
}
