# App HRD Golang

Aplikasi manajemen karyawan sederhana yang dibangun menggunakan Go (Golang) dengan database MySQL. Aplikasi ini menyediakan fungsi CRUD (Create, Read, Update, Delete) untuk mengelola data karyawan.

## Fitur

- **Dashboard Karyawan**: Menampilkan daftar semua karyawan
- **Tambah Karyawan**: Formulir untuk menambahkan karyawan baru
- **Edit Karyawan**: Formulir untuk mengubah data karyawan yang sudah ada
- **Hapus Karyawan**: Menghapus data karyawan dari database
- **Interface Web**: Menggunakan Bootstrap untuk tampilan yang responsif

## Struktur Project

```
├── main.go                 # Entry point aplikasi
├── go.mod                 # Go module dependencies
├── go.sum                 # Go dependencies checksum
├── controllers/           # Handler untuk HTTP requests
│   ├── ConnServer.go      # Handler untuk root endpoint
│   ├── create_karyawan.go # Handler untuk menambah karyawan
│   ├── delete_karyawan.go # Handler untuk menghapus karyawan
│   ├── index_karyawan.go  # Handler untuk menampilkan daftar karyawan
│   └── update_karyawan.go # Handler untuk mengubah data karyawan
├── database/              # Konfigurasi database
│   └── database.go        # Inisialisasi koneksi database
├── routes/                # Route mapping
│   └── route.go           # Definisi semua routes
└── views/                 # Template HTML
    ├── create.html        # Form tambah karyawan
    ├── index.html         # Daftar karyawan
    └── update.html        # Form edit karyawan
```

## Requirements

- Go 1.25.1 atau lebih tinggi
- MySQL Server
- Database dengan nama `db_hrd_golang`

## Dependencies

- `github.com/go-sql-driver/mysql` - MySQL driver untuk Go

## Setup Database

Buat database MySQL dengan nama `db_hrd_golang` dan tabel `employee`:

```sql
CREATE DATABASE db_hrd_golang;

USE db_hrd_golang;

CREATE TABLE employee (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    npwp VARCHAR(255) NOT NULL,
    alamat TEXT NOT NULL
);
```

## Konfigurasi Database

Aplikasi menggunakan konfigurasi database berikut (dapat diubah di [`database/database.go`](database/database.go)):
- Host: localhost
- Port: 3306
- Database: db_hrd_golang
- User: root
- Password: (kosong)

## Instalasi & Menjalankan Aplikasi

1. Clone atau download project ini
2. Pastikan MySQL server berjalan dan database sudah dibuat
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Jalankan aplikasi:
   ```bash
   go run main.go
   ```
5. Akses aplikasi di browser: `http://localhost:8080`

## Endpoints

| Method | Endpoint | Fungsi |
|--------|----------|---------|
| GET | `/` | Halaman utama |
| GET | `/karyawan` | Menampilkan daftar karyawan |
| GET | `/karyawan/create` | Form tambah karyawan |
| POST | `/karyawan/create` | Proses tambah karyawan |
| GET | `/karyawan/update?id={id}` | Form edit karyawan |
| POST | `/karyawan/update?id={id}` | Proses update karyawan |
| GET | `/karyawan/delete?id={id}` | Hapus karyawan |

## Data Model

Struktur data karyawan ([`controllers/index_karyawan.go`](controllers/index_karyawan.go)):

```go
type Karyawan struct {
    Id     string
    Nama   string
    NPWP   string
    Alamat string
}
```

## Teknologi yang Digunakan

- **Backend**: Go (Golang)
- **Database**: MySQL
- **Frontend**: HTML, Bootstrap 5
- **Template Engine**: Go HTML Template
- **HTTP Router**: Go standard library (net/http)

## File Penting

- [`main.go`](main.go) - Entry point aplikasi, mengatur server dan routing
- [`database/database.go`](database/database.go) - Konfigurasi dan inisialisasi database
- [`routes/route.go`](routes/route.go) - Mapping semua routes ke controller
- [`controllers/`](controllers/) - Berisi semua handler untuk HTTP requests
- [`views/`](views/) - Template HTML untuk interface pengguna

## Pengembangan

Untuk menambahkan fitur baru:
1. Buat handler baru di folder `controllers/`
2. Tambahkan route baru di [`routes/route.go`](routes/route.go)
3. Buat template HTML jika diperlukan di folder `views/`

## Kontribusi

Silakan buat pull request untuk perbaikan atau penambahan fitur baru.