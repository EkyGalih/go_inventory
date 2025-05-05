package config

import (
	"inventaris/entities"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "modernc.org/sqlite" // Gunakan modernc.org/sqlite agar tidak butuh CGO
)

var DB *gorm.DB

// ConnectDB melakukan koneksi ke database SQLite
func ConnectDB() {
	// Path database
	dbDir := "./database"
	dbPath := dbDir + "/inventaris.db"

	// Cek apakah folder database ada, jika tidak buat
	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		if err := os.MkdirAll(dbDir, os.ModePerm); err != nil {
			log.Fatalf("Gagal membuat folder database: %v", err)
		}
		log.Println("Folder database berhasil dibuat")
	}

	// Cek apakah file database ada, jika tidak buat file kosong
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		file, err := os.Create(dbPath)
		if err != nil {
			log.Fatalf("Gagal membuat file database: %v", err)
		}
		file.Close()
		log.Println("File database berhasil dibuat")
	}

	// Gunakan modernc.org/sqlite agar tidak butuh CGO
	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal koneksi ke database: %v", err)
	}

	log.Println("Database berhasil terkoneksi")
}

// InitDB melakukan migrasi database
func InitDB() {
	// Pastikan DB tidak nil sebelum melakukan migrasi
	if DB == nil {
		log.Fatal("Database belum dikoneksikan! Panggil ConnectDB() terlebih dahulu.")
	}

	// Lakukan migrasi
	err := DB.AutoMigrate(
		&entities.Aset{},
		&entities.Distribusi{},
		&entities.Pegawai{},
		&entities.Category{},
		&entities.Tipe{},
	)
	if err != nil {
		log.Fatalf("Gagal melakukan migrasi database: %v", err)
	}

	log.Println("Migrasi database berhasil")
}
