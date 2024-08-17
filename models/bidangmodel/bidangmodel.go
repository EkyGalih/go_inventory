package bidangmodel

import (
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