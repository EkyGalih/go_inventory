package entities

import "time"

type Category struct {
	Id            string
	Nama_Kategori string
	Created_At    time.Time
	Updated_At    time.Time
}
