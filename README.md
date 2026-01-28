# Tugas 02 - Bootcamp Jago Golang Basic with Umam 🚀

[![Golang](https://img.shields.io/badge/Golang-1.25%2B-blue.svg)](https://golang.org/)
[![Supabase](https://img.shields.io/badge/Supabase-green)](https://www.supabase.com/)
---

## 🧐 Tugas Session 2 Bootcamp Jago Golang Basic with Umam untuk membuat RESTful API untuk menu products & categories.

Dalam repositori ini kita menerapkan `Golang` sebagai platform dasar bahasa pemrograman yang digunakan dalam pembuatan `Rest API`.
Di dalam repositori ini juga kami terapkan `Supabase` sebagai media penyimpanan untuk mempermudah dalam pengerjaan di ranah menejemen database PostgreSQL.

### 🛠️ Dibangun Dengan (The Tech Stack)

Proyek ini dikembangkan menggunakan teknologi-teknologi utama berikut:

* [![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/) 
* [![Fiber](https://img.shields.io/badge/Fiber-v2-%2300ADEE.svg?style=for-the-badge&logo=gofiber&logoColor=white)](https://docs.gofiber.io)
* [![Supabase](https://img.shields.io/badge/supabase-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)](https://www.postgresql.org/)

---

## 🏁 Memulai (Getting Started)

Bagian ini memandu Anda untuk menyiapkan dan menjalankan proyek di lingkungan lokal Anda untuk tujuan pengembangan dan pengujian.

### ⚙️ Prerequisites (Prasyarat)

Pastikan Anda telah menginstal yang berikut ini:

* **Golang** (Versi 1.25 atau lebih tinggi)
* **PostgreSQL** (Database)
* **Git**
* **Fiber v2**

### 📦 Installation (Instalasi)

1.  **Clone** repositori ini:
    ```bash
     git clone git@github.com:heru-oktafian/tugas-01-bootacmp.git
     cd cores
     go mod init "github.com/heru-oktafian/tugas-01-bootacmp"
    ```

2.  **Siapkan Database:**
    * Buat database PostgreSQL baru.
    * Konfigurasi koneksi database Anda di file `.env` dengan menjadikan acuan `.example_env`.

3.  **Siapkan Environment (Lingkungan):**
    * Duplikasi file `.example_env` dan ganti namanya menjadi `.env`.
    * Isi variabel-variabel yang diperlukan (`DB_HOST`, `DB_USER`, `REDIS_HOST`, `JWT_SECRET`, dll.).

4.  **Jalankan Proyek:**
    ```bash
    go run main.go
    # Atau gunakan: go build && ./[nama executable]
    ```

Proyek akan berjalan di `http://localhost:9001`.
---

## 🛠️ Tahapan Pembuatan

Dalam repository ini, kami juga sertakan proses serta tahapan dalam pembuatannya, serta aspek yang terdapat di dalamnya apa saja.

### Endpoint Utama

| Methode | Endpoints | Deskripsi |
| :--- | :--- | :--- |
| Get | `/categories` | Tampilkan semua kategori. |
| Post | `/categories` | Create satu kategori baru. |
| Put | `/categories/{id}` | Update data kategori sesuai ID. |
| Get | `/categories/{id}` | Tampilkan kategori sesuai ID. |
| Delete | `/categories/{id}` | Hapus kategori sesuai ID. |
---

## ✉️ Kontak (Contact)

Heru Oktafian, ST., CTT - [@heru-oktafian](https://x.com/HeruOktafianST) - [info@heruoktafian.com](mailto:info@heruoktafian.com)

Tautan Proyek: [https://github.com/heru-oktafian/tugas-01-bootcamp](https://github.com/heru-oktafian/tugas-01-bootcamp)
