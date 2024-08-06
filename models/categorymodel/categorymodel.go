package categorymodel

import (
	"inventaris/config"
	"inventaris/entities"

	"github.com/google/uuid"
)

func GetAll() []entities.Category {
	rows, err := config.DB.Query(`SELECT * FROM kategori_aset`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.Id, &category.Nama_Kategori, &category.Created_At, &category.Updated_At); err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}

	return categories
}

func Create(category entities.Category) (bool, error) {
	newUUID := uuid.New()
	result, err := config.DB.Exec(`
	INSERT INTO kategori_aset (id, nama_kategori, created_at, updated_at)
	VALUES (?, ?, ?, ?)`,
		newUUID,
		category.Nama_Kategori,
		category.Created_At,
		category.Updated_At,
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

func Detail(id int) entities.Category {
	row := config.DB.QueryRow(`SELECT * FROM kategori_aset WHERE id = ?`, id)

	var category entities.Category
	if err := row.Scan(&category.Id, &category.Nama_Kategori, &category.Created_At, &category.Updated_At); err != nil {
		panic(err.Error())
	}

	return category
}

func Update(id int, category entities.Category) bool {
	query, err := config.DB.Exec(`UPDATE categories SET nama_kategori = ?, updated_at = ? WHERE id = ?`, category.Nama_Kategori, category.Updated_At, id)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec(`DELETE FROM categories WHERE id = ?`, id)
	return err
}
