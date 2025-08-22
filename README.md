# ShoppinGO: Simple E-Commerce API (Go Backend Developer)

Proyek ini adalah implementasi backend API untuk aplikasi e-commerce sederhana, yang dibuat sebagai bagian dari LnT Camp 2025 oleh BNCC. API ini melayani dua bagian utama aplikasi: halaman landing publik untuk menampilkan produk dan panel admin untuk manajemen data.


---

## Teknologi yang Digunakan

-   **Backend:** Go (Golang)
-   **Web Framework:** Gin
-   **ORM:** GORM
-   **Database:** MySQL
-   **Autentikasi:** JWT (JSON Web Tokens)
-   **Export:** Excelize for Go
-   **Frontend:** TypeScript, Vite

---

## Fitur Utama

-   **Public API:** Menampilkan produk terbaru dan produk yang tersedia untuk semua pengunjung.
-   **Admin API:**
    -   Autentikasi admin yang aman menggunakan JWT.
    -   Dasbor yang menampilkan statistik kunci (jumlah user, produk, dll.).
    -   Fungsionalitas CRUD (Create, Read, Update, Delete) penuh untuk data User dan Produk.
    -   Fitur untuk mengekspor data produk ke dalam format file Excel.

---

## Struktur API Endpoint

### Public API

| Method | Endpoint              | Deskripsi                                        |
| :----- | :-------------------- | :----------------------------------------------- |
| `GET`  | `/products/latest`    | Menampilkan daftar produk terbaru.               |
| `GET`  | `/products/available` | Menampilkan daftar produk yang stoknya lebih dari 0. |

### Admin API

| Method         | Endpoint                 | Deskripsi                                                                    |
| :------------- | :----------------------- | :--------------------------------------------------------------------------- |
| `POST`         | `/admin/login`           | Login admin dan mendapatkan token JWT.                                       |
| `GET`          | `/admin/dashboard`       | Menampilkan data statistik untuk dasbor admin.                               |
| `GET, POST, PUT` | `/admin/users`           | Fungsionalitas CRUD untuk data user (membutuhkan autentikasi).             |
| `GET, POST, PUT` | `/admin/products`        | Fungsionalitas CRUD untuk data produk (membutuhkan autentikasi).             |
| `GET`          | `/admin/products/export` | Mengunduh data semua produk dalam format `.xlsx` (membutuhkan autentikasi). |

---

## Cara Menjalankan Proyek

### Prasyarat

-   Go (versi 1.20+) terinstal.
-   Node.js dan npm terinstal.
-   MySQL server berjalan.

### Backend Setup

1.  **Clone repository ini.**
2.  Masuk ke direktori `backend`: `cd backend`.
3.  Buat file `.env` dari contoh yang ada (`.env.example` jika ada) dan sesuaikan koneksi database Anda.
    ```env
    DB_USER=root
    DB_PASS=password_anda
    DB_HOST=127.0.0.1
    DB_PORT=3306
    DB_NAME=ecommerce_db
    JWT_SECRET_KEY=kunci_rahasia_anda
    ```
4.  Instal dependensi Go: `go mod tidy`.
5.  Jalankan server backend: `go run main.go`. Server akan berjalan di `http://localhost:8080`.

### Frontend Setup

1.  Buka terminal baru dan masuk ke direktori `frontend`: `cd frontend`.
2.  Instal dependensi Node.js: `npm install`.
3.  Jalankan server pengembangan Vite: `npm run dev`. Aplikasi akan berjalan di `http://localhost:5173`.

⚠️ The frontend is still under development.

---

## Dokumentasi API (Postman)

Koleksi Postman yang berisi semua endpoint API dan contoh penggunaannya dapat diakses melalui link berikut:

https://documenter.getpostman.com/view/45041293/2sB3BHnUjZ)

