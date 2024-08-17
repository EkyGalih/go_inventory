package pegawaimodel

import (
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