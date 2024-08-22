package pegawaimodel

import (
	"database/sql"
	"fmt"
	"inventaris/config"
	"inventaris/entities"
)

func GetALl() []entities.Pegawai {
	rows, err := config.DB.Query(
	`SELECT 
		id, name, nip
	FROM
		pegawai
	`)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var pegawais []entities.Pegawai

	for rows.Next() {
		var pegawai entities.Pegawai

		if err := rows.Scan(
			&pegawai.Id,
			&pegawai.Name,
			&pegawai.Nip,
		); err != nil {
			panic(err.Error())
		}

		pegawais = append(pegawais, pegawai)
	}

	return pegawais
}

func Detail(id string) (entities.Pegawai, error) {
	row := config.DB.QueryRow(`SELECT id, name, nip FROM pegawai WHERE id = ?`, id)

	var pegawai entities.Pegawai
	if err := row.Scan(&pegawai.Id, &pegawai.Name, &pegawai.Nip); err != nil {
		if err == sql.ErrNoRows {
			return pegawai, fmt.Errorf("pegawai not found with id %s", id)
		}
		return pegawai, fmt.Errorf("failed to retrieve pegawai: %w", err)
	}

	return pegawai, nil
}