# ShoppinGO: Simple E-Commerce API & Frontend

Proyek ini adalah implementasi *full-stack* untuk aplikasi e-commerce sederhana, yang dibuat sebagai bagian dari LnT Camp 2025 oleh BNCC. Aplikasi ini terdiri dari dua bagian utama: **Backend API** yang andal dibuat dengan Go, dan **Frontend interaktif** yang modern dibuat dengan TypeScript dan Vite.


---

## Teknologi yang Digunakan

| Kategori | Teknologi | Deskripsi |
| :--- | :--- | :--- |
| **Backend** | Go (Golang) | Bahasa utama untuk membangun API yang cepat dan efisien. |
| | Gin | Framework web minimalis untuk routing dan middleware. |
| | GORM | ORM yang kuat untuk interaksi dengan database. |
| | MySQL | Sistem manajemen database relasional. |
| | JWT | Standar untuk autentikasi dan otorisasi yang aman. |
| | Excelize | Library untuk membuat dan memanipulasi file Excel (.xlsx). |
| **Frontend** | TypeScript | Superset JavaScript yang menambahkan tipe statis untuk kode yang lebih aman. |
| | Vite | Alat pengembangan frontend modern yang sangat cepat. |
| | Swiper.js | Library carousel modern untuk slider interaktif. |

---

## Fitur Utama

### - Sisi Publik (Landing Page)
- **Tampilan Produk:** Menampilkan daftar produk terbaru dalam bentuk *slider* dan semua produk yang tersedia dalam bentuk *grid*.
- **Carousel Interaktif:** Banner promosi dinamis yang mendukung gambar dan video, dilengkapi navigasi dan auto-play.
- **Detail Produk (Pop-up):** Tampilan detail produk yang bersih dan modern saat kartu produk diklik.
- **Tombol Dinamis:** Tombol "MASUK" yang berubah menjadi "DASHBOARD" jika admin sudah login.

### - Sisi Admin (Dashboard)
- **Autentikasi Aman:** Sistem login yang aman menggunakan JSON Web Tokens (JWT).
- **Dashboard Statistik:** Menampilkan ringkasan data kunci seperti jumlah total user, user aktif, total produk, dan produk tersedia.
- **Manajemen Produk (CRUD):**
    - Membuat, membaca, mengedit data produk.
    - Upload gambar produk.
    - **Penanggung Jawab Otomatis:** Sistem secara otomatis mencatat admin yang terakhir kali membuat/mengedit produk.
- **Manajemen User (CRUD):**
    - Membuat, membaca, mengedit data pengguna.
    - Mengaktifkan atau menonaktifkan akun pengguna.
- **Ekspor ke Excel:** Fungsionalitas untuk mengunduh seluruh data produk dalam format file `.xlsx`.
- **UI Profesional:** Sidebar statis dan konten yang dapat di-scroll untuk pengalaman pengguna yang lebih baik.

---

## Struktur API Endpoint

#### Public API
| Method | Endpoint | Deskripsi |
| :--- | :--- | :--- |
| `GET` | `/products/latest` | Menampilkan 5 produk terbaru. |
| `GET` | `/products/available` | Menampilkan produk yang stoknya lebih dari 0. |
| `GET` | `/products/:id` | Mengambil detail satu produk berdasarkan ID. |

#### Admin API (Membutuhkan Token JWT)
| Method | Endpoint | Deskripsi |
| :--- | :--- | :--- |
| `POST` | `/admin/login` | Login sebagai admin untuk mendapatkan token. |
| `GET` | `/admin/dashboard` | Mengambil data statistik untuk dasbor. |
| `GET`, `POST`, `PUT` | `/admin/users` | Fungsionalitas CRUD untuk data user. |
| `GET`, `POST`, `PUT` | `/admin/products` | Fungsionalitas CRUD untuk data produk. |
| `POST` | `/admin/products/:id/upload`| Mengunggah gambar untuk produk tertentu. |
| `GET` | `/admin/products/export` | Mengunduh data produk dalam format .xlsx. |

---

## Cara Menjalankan Proyek

### Prasyarat
- [Go](https://golang.org/) (versi 1.20+) terinstal.
- [Node.js](https://nodejs.org/) (versi 18+) dan npm terinstal.
- Server MySQL berjalan.

### - Backend Setup
1.  Clone repositori ini: `git clone [URL_REPOSITORI_ANDA]`
2.  Masuk ke direktori backend: `cd backend`
3.  Buat file `.env` dan sesuaikan dengan konfigurasi database Anda:
    ```env
    DB_USER=root
    DB_PASS=password_anda
    DB_HOST=127.0.0.1
    DB_PORT=3306
    DB_NAME=ecommerce_db
    JWT_SECRET_KEY=kunci_rahasia_yang_sangat_aman
    ```
4.  Instal dependensi Go: `go mod tidy`
5.  Jalankan server backend: `go run main.go`
    > Server akan berjalan di `http://localhost:8080` dan secara otomatis melakukan *seeding* data awal.

### - Frontend Setup
1.  Buka terminal baru dan masuk ke direktori frontend: `cd frontend`
2.  Instal dependensi Node.js: `npm install`
3.  Jalankan server pengembangan Vite: `npm run dev`
    > Aplikasi akan dapat diakses di `http://localhost:5173`.

---

## Dokumentasi API (Postman)

Koleksi Postman yang berisi semua endpoint API beserta contoh penggunaannya dapat diakses melalui link berikut:

[**Lihat Dokumentasi API di Postman**](https://documenter.getpostman.com/view/45041293/2sB3BHnUjZ)
