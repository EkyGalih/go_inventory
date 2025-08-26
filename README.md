# 📦 Aplikasi Inventarisasi Aset

Proyek ini adalah aplikasi inventarisasi aset berbasis **Golang** yang dibangun menggunakan framework **Gin**.  
Aplikasi ini digunakan untuk mencatat, mengelola, dan menampilkan data aset dengan arsitektur **MVC (Model-View-Controller)**.

---

## 🚀 Fitur Utama
- Autentikasi menggunakan JWT (JSON Web Token)
- CRUD (Create, Read, Update, Delete) untuk data aset
- Penyimpanan data ke dalam database (contoh: PostgreSQL / MySQL / SQLite)
- Template rendering menggunakan HTML (dengan `helpers.go`)
- Struktur proyek terorganisir dengan baik (entities, models, controllers, routes)
- API endpoint yang bisa diintegrasikan dengan aplikasi lain

---

## 📂 Struktur Folder
```
.
├── controllers/       # Logika utama aplikasi (handle request & response)
├── entities/          # Definisi struktur tabel (model entitas)
├── models/            # Interaksi dengan database (query, exec, dll)
├── routes/            # Definisi rute aplikasi (API & Web)
├── templates/         # File HTML (frontend)
│   ├── layouts/       # Template utama (header, footer, navbar)
│   └── pages/         # Halaman konten (dashboard, aset, dll)
├── static/            # File statis (CSS, JS, Images)
├── utils/             # Helper function & konfigurasi tambahan
├── main.go            # Entry point aplikasi
├── go.mod             # Modul Go
└── README.md          # Dokumentasi proyek
```

---

## ⚙️ Instalasi & Menjalankan

### 1. Clone Repository
```bash
git clone https://github.com/username/inventarisasi-aset.git
cd inventarisasi-aset
```

### 2. Install Dependency
```bash
go mod tidy
```

### 3. Konfigurasi Database
Edit file `.env` untuk menyesuaikan konfigurasi database:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=password
DB_NAME=inventarisasi
JWT_SECRET=your-secret-key
```

### 4. Jalankan Aplikasi
```bash
air
```
atau tanpa hot reload:
```bash
go run main.go
```

---

## 🔑 Autentikasi JWT
- Login menghasilkan **Access Token** (JWT)
- Token disimpan di header `Authorization: Bearer <token>`
- Semua endpoint yang membutuhkan autentikasi wajib mengirimkan token valid

---

## 📌 Endpoint Utama

### Auth
- `POST /login` → Login user
- `POST /register` → Registrasi user baru

### Aset
- `GET /aset` → Menampilkan semua aset
- `POST /aset` → Tambah aset baru
- `PUT /aset/:id` → Update data aset
- `DELETE /aset/:id` → Hapus aset

---

## 🖥️ Teknologi yang Digunakan
- [Golang](https://golang.org/)
- [Gin Gonic](https://gin-gonic.com/)
- [GORM](https://gorm.io/) (ORM untuk database)
- [JWT](https://jwt.io/)
- [Air](https://github.com/cosmtrek/air) (Live reload)
- HTML Template

---

## 👨‍💻 Author
Dibuat oleh **Eky Galih Gunanda**  
📧 Email: [your.email@example.com](mailto:your.email@example.com)

---

## 📜 Lisensi
Proyek ini dirilis dengan lisensi **MIT**.  
Silakan gunakan, modifikasi, dan distribusikan sesuai kebutuhan Anda.
