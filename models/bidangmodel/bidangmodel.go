package bidangmodel

import (
	"database/sql"
	"fmt"
	"inventaris/config"
	"inventaris/entities"
)

func GetAll() []entities.Bidang {
	rows, err := config.DB.Query(
		`SELECT 
			id, 
			nama_bidang
		FROM 
			bidang;
		`)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var bidangs []entities.Bidang

	for rows.Next() {
		var bidang entities.Bidang

		if err := rows.Scan(
			&bidang.Id,
			&bidang.Nama_Bidang,
		); err != nil {
			panic(err.Error())
		}

		bidangs = append(bidangs, bidang)
	}

	return bidangs
}


func Detail(id string) (entities.Bidang, error) {
	row := config.DB.QueryRow(`SELECT id, Nama_Bidang FROM bidang WHERE id = ?`, id)

	var bidang entities.Bidang
	if err := row.Scan(&bidang.Id, &bidang.Nama_Bidang); err != nil {
		if err == sql.ErrNoRows {
			return bidang, fmt.Errorf("bidang not found with id %s", id)
		}
		return bidang, fmt.Errorf("failed to retrieve bidang: %w", err)
	}

	return bidang, nil
}