package entities

import "time"

type Bidang struct {
	ID         string    `json:"ID"`
	NamaBidang string    `json:"NamaBidang"`
	CreatedAt  time.Time `json:"CreatedAt"`
	UpdatedAt  time.Time `json:"UpdatedAt"`
}
