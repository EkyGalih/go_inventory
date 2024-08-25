package entities

type User struct {
	Id        string `gorm:"primaryKey"`
	Nama_User string
	Email     string
	Username  string
	Password  string
}
