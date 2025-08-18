# Backend API - ShoppinGO Mini Project

Direktori ini berisi layanan backend lengkap untuk aplikasi E-Commerce ShoppinGO. Dibangun dengan Go (Golang) dan Gin, API ini menyediakan semua endpoint yang diperlukan untuk manajemen user, produk, autentikasi, dan analisis data melalui dasbor.

Proyek ini dibuat sebagai bagian dari **LnT Camp 2025** oleh **BNCC**.

---

## Tech Stack

-   **Bahasa Pemrograman:** Go (Golang)
-   **Web Framework:** Gin
-   **ORM:** GORM
-   **Database:** MySQL
-   **Autentikasi:** JWT (JSON Web Tokens)
-   **Fitur Export:** Excelize for Go

---

## Fitur Utama

-   **API Publik:** Menyediakan data produk yang dapat diakses secara umum untuk halaman landing.
-   **API Admin:** Kumpulan endpoint yang dilindungi untuk melakukan manajemen aplikasi, mencakup:
    -   Login admin yang aman menggunakan JWT.
    -   Dasbor statistik (total user, produk, dll.).
    -   CRUD (Create, Read, Update, Delete) penuh untuk data User.
    -   CRUD penuh untuk data Produk, termasuk fungsionalitas upload gambar.
    -   Fitur ekspor data produk ke file Excel (`.xlsx`).
-   **Database:** Menggunakan GORM AutoMigrate untuk membuat skema tabel secara otomatis dan Seeder untuk mengisi data awal.

---

## Struktur API Endpoint

### Public API

| Method | Endpoint              | Deskripsi                                        |
| :----- | :-------------------- | :----------------------------------------------- |
| `GET`  | `/products/latest`    | Menampilkan daftar produk terbaru.               |
| `GET`  | `/products/available` | Menampilkan daftar produk yang stoknya tersedia. |

### Admin API (Membutuhkan Token Autentikasi)

| Method         | Endpoint                 | Deskripsi                                          |
| :------------- | :----------------------- | :------------------------------------------------- |
| `POST`         | `/admin/login`           | Login sebagai admin untuk mendapatkan token JWT.     |
| `GET`          | `/admin/dashboard`       | Mengambil data statistik agregat untuk dasbor.     |
| `GET, POST, PUT` | `/admin/users`           | Fungsionalitas CRUD untuk data user.               |
| `GET, POST, PUT` | `/admin/products`        | Fungsionalitas CRUD untuk data produk.             |
| `POST`         | `/admin/products/:id/upload` | Upload gambar untuk produk tertentu.               |
| `GET`          | `/admin/products/export` | Mengunduh laporan data produk dalam format Excel. |

---

## Panduan Setup & Instalasi

### 1. Prasyarat
-   Go (versi 1.20 atau lebih baru)
-   MySQL Server

### 2. Konfigurasi Database
-   Buat sebuah database baru di MySQL Anda, misalnya dengan nama `shoppin_go_db`.
-   Di dalam direktori `backend/`, buat sebuah file bernama `.env`.
-   Salin konten di bawah ini ke dalam file `.env` Anda dan sesuaikan dengan konfigurasi MySQL Anda.

    ```env
    # Konfigurasi Database
    DB_USER=root
    DB_PASS=password_mysql_anda
    DB_HOST=127.0.0.1
    DB_PORT=3306
    DB_NAME=shoppin_go_db

    # Kunci Rahasia untuk JWT
    JWT_SECRET_KEY=ini_adalah_kunci_rahasia_yang_sangat_aman
    ```

### 3. Instalasi & Menjalankan Server
1.  Buka terminal dan navigasikan ke direktori `backend/`.
2.  Instal semua dependensi yang dibutuhkan:
    ```bash
    go mod tidy
    ```
3.  Jalankan server:
    ```bash
    go run main.go
    ```
-   Saat pertama kali dijalankan, GORM AutoMigrate akan secara otomatis membuat semua tabel yang diperlukan di database Anda.
-   Seeder juga akan berjalan untuk mengisi data awal (user admin dan beberapa produk) agar aplikasi siap untuk diuji.
-   Server akan berjalan di `http://localhost:8080`.

---

## Dokumentasi API (Postman)

Dokumentasi lengkap untuk semua endpoint, termasuk contoh request dan response, telah dipublikasikan dan dapat diakses melalui link di bawah ini:

[**Lihat Dokumentasi API Lengkap di Postman**](https://documenter.getpostman.com/view/45041293/2sB3BHnUjZ)
