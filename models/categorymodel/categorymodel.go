package categorymodel

import (
	"database/sql"
	"fmt"
	"inventaris/config"
	"inventaris/entities"

	"github.com/google/uuid"
)

func GetAll() []entities.Category {
	rows, err := config.DB.Query(`SELECT * FROM kategori_aset ORDER BY updated_at DESC`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.Id, &category.Nama_Kategori, &category.Deskripsi, &category.Created_At, &category.Updated_At); err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}

	return categories
}

func Create(category entities.Category) (bool, error) {
	newUUID := uuid.New()
	result, err := config.DB.Exec(`
	INSERT INTO kategori_aset (id, nama_kategori, deskripsi, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?)`,
		newUUID,
		category.Nama_Kategori,
		category.Deskripsi,
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

func Detail(id string) (entities.Category, error) {
	row := config.DB.QueryRow(`SELECT * FROM kategori_aset WHERE id = ?`, id)

	var category entities.Category
	if err := row.Scan(&category.Id, &category.Nama_Kategori, &category.Deskripsi, &category.Created_At, &category.Updated_At); err != nil {
		if err == sql.ErrNoRows {
			return category, fmt.Errorf("no category found with id %s", id)
		}
		return category, fmt.Errorf("failed to retrieve category: %w", err)
	}

	return category, nil
}
func Update(id string, category entities.Category) bool {
	query, err := config.DB.Exec(`UPDATE kategori_aset SET nama_kategori = ?, deskripsi = ?, updated_at = ? WHERE id = ?`, category.Nama_Kategori, category.Deskripsi, category.Updated_At, id)
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
	_, err := config.DB.Exec(`DELETE FROM kategori_aset WHERE id = ?`, id)
	return err
}
