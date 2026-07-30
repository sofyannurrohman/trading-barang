---
name: erp-pkp-development
description: Panduan pengembangan Sistem Manajemen Perdagangan & ERP Mini (PKP) berdasarkan arsitektur Vue.js, Go + Gin, dan PostgreSQL.
---

# Panduan Pengembangan ERP PKP

Skill ini memberikan panduan arsitektur dan praktik terbaik dalam pengembangan Sistem Manajemen Perdagangan & ERP Mini (PKP), sesuai dengan dokumen `technical-development-plan-erp-pkp.md`.

## 1. Tech Stack Utama
- **Frontend**: Vue 3 (Composition API), Vite, Pinia, Vue Router, Tailwind CSS (khususnya untuk print layout).
- **Backend**: Golang dengan Gin Framework.
- **Database**: PostgreSQL.
- **Komunikasi**: REST API dengan format JSON. Axios di sisi frontend.
- **Autentikasi**: JWT (Access & Refresh Token).

## 2. Aturan Backend (Go + Gin)
- **Arsitektur Layered**: Wajib memisahkan logika ke dalam layer:
  - `Handler` (Controller API, parsing request)
  - `Service` (Logika bisnis, HPP, pajak)
  - `Repository` (Query database menggunakan GORM atau SQLx)
- **Presisi Uang**: **DILARANG KERAS** menggunakan `float64` untuk data uang/harga. Wajib menggunakan `shopspring/decimal` di Go dan dipetakan ke tipe `DECIMAL` di PostgreSQL untuk menghindari *floating point rounding error*.
- **Database Transaction & Locking**: Operasi kritikal seperti pengurangan/penambahan stok (`Stock-In`, `Sales`) harus dibungkus dalam satu DB Transaction dengan row-level lock (`SELECT ... FOR UPDATE`) untuk mencegah *race condition*.
- **Nomor Faktur**: Generate nomor faktur menggunakan PostgreSQL sequence agar *concurrent-safe*.

## 3. Aturan Database (PostgreSQL)
- **Tipe Data**: Gunakan `UUID` untuk primary key tabel utama (bisa menggunakan `gen_random_uuid()`).
- **Audit Trail**: Selalu sertakan `created_at`, `updated_at`, dan `created_by` (jika relevan).
- **Relasi & Indexing**: Buat index pada kolom pencarian (seperti tipe partner, SKU, status transaksi) untuk performa.
- **Soft Delete**: Pertimbangkan penggunaan is_active atau status seperti VOID untuk transaksi daripada menghapus data (menjaga rekam jejak audit).

## 4. Aturan Frontend (Vue.js)
- **State Management**: Gunakan Pinia untuk menyimpan state global (Auth, Produk, Penjualan).
- **Form Validation**: Gunakan `vee-validate` bersama `yup` untuk validasi sisi klien.
- **Print Layout**: Untuk halaman cetak faktur A4, buat layout terpisah (`PrintLayout.vue`) tanpa sidebar/navbar, dengan konfigurasi CSS khusus cetak (`@page { size: A4; margin: 15mm }`).

## 5. Logika Bisnis Kritis
- **HPP (Harga Pokok Penjualan)**: Menggunakan metode **Moving Average**. Kalkulasi dilakukan otomatis saat ada barang masuk (Stock-In).
- **Pajak (PPN)**: Konfigurasi pajak disimpan di tabel `tax_configs` agar tidak hardcode. Kalkulasi PPN harus mendukung opsi *tax include* (harga sudah termasuk pajak) dan *tax exclude* (pajak ditambahkan di luar harga).

## 6. Alur Testing
- **Backend**: Wajib menambahkan unit test untuk `tax_service` dan `inventory_service` (table-driven test).
- **Frontend**: Component test untuk form, dan E2E test untuk flow kritis seperti pembuatan faktur.
