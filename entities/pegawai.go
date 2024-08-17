package entities

import "database/sql"

type Pegawai struct {
	Id   string `gorm:"primaryKey"`
	Name string
	Nip  sql.NullString
}
