# Technical Development Plan: Sistem Manajemen Perdagangan & ERP Mini (PKP)

**Tech Stack:** Vue.js (Frontend) · Golang + Gin (Backend) · PostgreSQL (Database)

---

## 1. Ringkasan Arsitektur

```
┌─────────────────┐        REST API (JSON)        ┌──────────────────┐        ┌──────────────┐
│   Vue.js SPA     │  ───────────────────────────▶ │   Go + Gin API    │ ──────▶│ PostgreSQL   │
│  (Vite, Pinia,   │ ◀─────────────────────────── │  (Clean/Layered)  │ ◀──────│  (ACID)      │
│  Vue Router)     │                               │                   │        └──────────────┘
└─────────────────┘                               └──────────────────┘
```

- **Pola arsitektur backend:** Layered architecture (Handler → Service → Repository) agar logika bisnis (HPP, pajak) terisolasi dari detail HTTP dan SQL.
- **Pola frontend:** SPA dengan Vue 3 Composition API, state management Pinia, komunikasi via Axios/Fetch ke REST API.
- **Autentikasi:** JWT (access token + refresh token), middleware Gin untuk proteksi route.
- **Deployment:** Docker Compose (service: `api`, `web`, `db`, `nginx` sebagai reverse proxy).

---

## 2. Struktur Proyek

### 2.1. Backend (Go + Gin)

```
backend/
├── cmd/
│   └── api/
│       └── main.go                 # entry point, load config, start server
├── internal/
│   ├── config/                     # load .env, konfigurasi DB, JWT secret
│   ├── domain/                     # struct entity murni (Product, Partner, dsb)
│   ├── dto/                        # request/response payload (input/output API)
│   ├── handler/                    # Gin handler (controller layer)
│   │   ├── product_handler.go
│   │   ├── partner_handler.go
│   │   ├── inventory_handler.go
│   │   ├── sales_handler.go
│   │   └── auth_handler.go
│   ├── service/                    # business logic (HPP, pajak, validasi)
│   │   ├── product_service.go
│   │   ├── inventory_service.go
│   │   ├── sales_service.go
│   │   └── tax_service.go
│   ├── repository/                 # akses database (SQL/GORM query)
│   │   ├── product_repo.go
│   │   ├── partner_repo.go
│   │   ├── inventory_repo.go
│   │   └── sales_repo.go
│   ├── middleware/                 # JWT auth, logging, CORS, error handler
│   └── router/                     # deklarasi route Gin
├── migrations/                     # file migrasi SQL (golang-migrate)
├── pkg/
│   └── utils/                      # helper: generate invoice number, formatter
├── go.mod
└── Dockerfile
```

**Library kunci:**
| Kebutuhan | Library |
| :--- | :--- |
| Router & middleware | `gin-gonic/gin` |
| ORM / query builder | `gorm.io/gorm` + driver `gorm.io/driver/postgres` (atau `jmoiron/sqlx` jika ingin kontrol SQL penuh — direkomendasikan untuk transaksi finansial yang presisi) |
| Migrasi DB | `golang-migrate/migrate` |
| Validasi input | `go-playground/validator` |
| JWT | `golang-jwt/jwt` |
| Konfigurasi | `spf13/viper` atau `joho/godotenv` |
| Logging | `uber-go/zap` |
| UUID | `google/uuid` |
| Decimal precision (WAJIB, jangan pakai float untuk uang) | `shopspring/decimal` |

> **Catatan penting:** Semua kolom `DECIMAL` (harga, HPP, pajak) harus dipetakan ke tipe `decimal.Decimal` dari `shopspring/decimal` di Go, bukan `float64`, untuk menghindari *floating point rounding error* pada perhitungan uang.

### 2.2. Frontend (Vue.js)

```
frontend/
├── src/
│   ├── api/                        # axios instance + service per modul
│   │   ├── axiosClient.js
│   │   ├── productApi.js
│   │   ├── partnerApi.js
│   │   ├── inventoryApi.js
│   │   └── salesApi.js
│   ├── assets/
│   ├── components/
│   │   ├── common/                 # DataTable, Modal, Pagination, Toast
│   │   └── forms/                  # ProductForm, PartnerForm, InvoiceForm
│   ├── layouts/
│   │   ├── AdminLayout.vue         # dengan sidebar/navbar
│   │   └── PrintLayout.vue         # kosong, khusus halaman cetak A4
│   ├── views/
│   │   ├── auth/LoginView.vue
│   │   ├── dashboard/DashboardView.vue
│   │   ├── products/ (List, Detail, Form)
│   │   ├── partners/ (List, Form)
│   │   ├── inventory/ (StockIn, StockAdjustment, Ledger)
│   │   ├── sales/ (InvoiceList, InvoiceForm)
│   │   └── invoice/PrintInvoiceView.vue   # route /invoice/print/:invoiceNo
│   ├── stores/                     # Pinia stores
│   │   ├── authStore.js
│   │   ├── productStore.js
│   │   └── salesStore.js
│   ├── router/index.js
│   ├── App.vue
│   └── main.js
├── vite.config.js
└── Dockerfile
```

**Library kunci:** `vue-router`, `pinia`, `axios`, `vee-validate` + `yup` (validasi form), `vue-toastification` (notifikasi), Tailwind CSS (styling cepat & konsisten untuk print layout).

---

## 3. Desain Skema Database (PostgreSQL) — Disempurnakan

Menambahkan beberapa penyesuaian dari draft awal: constraint, index, dan tabel tambahan untuk kelengkapan sistem (users, purchase orders, tax config).

```sql
-- Extension untuk UUID
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TYPE partner_type AS ENUM ('CUSTOMER', 'SUPPLIER');
CREATE TYPE trx_type AS ENUM ('IN', 'OUT', 'ADJUSTMENT');

-- 3.1 Users (untuk autentikasi & audit trail)
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role VARCHAR(20) NOT NULL DEFAULT 'staff', -- admin, staff
    created_at TIMESTAMPTZ DEFAULT now()
);

-- 3.2 Partners
CREATE TABLE partners (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(150) NOT NULL,
    type partner_type NOT NULL,
    npwp VARCHAR(30),
    nik VARCHAR(20),
    address TEXT,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);
CREATE INDEX idx_partners_type ON partners(type);

-- 3.3 Products
CREATE TABLE products (
    sku VARCHAR(50) PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    current_stock INT NOT NULL DEFAULT 0,
    current_hpp DECIMAL(18,2) NOT NULL DEFAULT 0,
    base_price DECIMAL(18,2) NOT NULL DEFAULT 0,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    CONSTRAINT chk_stock_non_negative CHECK (current_stock >= 0)
);

-- 3.4 Inventory Transactions (Buku Besar Stok)
CREATE TABLE inventory_transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    product_sku VARCHAR(50) NOT NULL REFERENCES products(sku),
    trx_type trx_type NOT NULL,
    qty INT NOT NULL,
    unit_cost DECIMAL(18,2) NOT NULL,
    reference_no VARCHAR(50),      -- link ke invoice_no / PO number
    note TEXT,
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMPTZ DEFAULT now()
);
CREATE INDEX idx_inv_trx_product ON inventory_transactions(product_sku);
CREATE INDEX idx_inv_trx_created ON inventory_transactions(created_at);

-- 3.5 Tax Configuration (agar tarif PPN tidak hardcode, mendukung perubahan 11% -> 12%)
CREATE TABLE tax_configs (
    id SERIAL PRIMARY KEY,
    rate DECIMAL(5,2) NOT NULL,       -- contoh: 11.00 atau 12.00
    effective_date DATE NOT NULL,
    is_active BOOLEAN DEFAULT true
);

-- 3.6 Sales Orders
CREATE TABLE sales_orders (
    invoice_no VARCHAR(50) PRIMARY KEY,
    partner_id UUID NOT NULL REFERENCES partners(id),
    total_dpp DECIMAL(18,2) NOT NULL,
    total_ppn DECIMAL(18,2) NOT NULL,
    grand_total DECIMAL(18,2) NOT NULL,
    tax_include BOOLEAN NOT NULL DEFAULT false,
    status VARCHAR(20) NOT NULL DEFAULT 'ISSUED', -- DRAFT, ISSUED, VOID
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMPTZ DEFAULT now()
);
CREATE INDEX idx_sales_partner ON sales_orders(partner_id);
CREATE INDEX idx_sales_created ON sales_orders(created_at);

-- 3.7 Sales Items
CREATE TABLE sales_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    invoice_no VARCHAR(50) NOT NULL REFERENCES sales_orders(invoice_no) ON DELETE CASCADE,
    product_sku VARCHAR(50) NOT NULL REFERENCES products(sku),
    qty INT NOT NULL,
    unit_hpp DECIMAL(18,2) NOT NULL,
    unit_dpp DECIMAL(18,2) NOT NULL,
    unit_ppn DECIMAL(18,2) NOT NULL,
    subtotal DECIMAL(18,2) NOT NULL
);
CREATE INDEX idx_sales_items_invoice ON sales_items(invoice_no);
```

**Catatan desain penting:**
- `inventory_transactions.reference_no` menghubungkan transaksi stok keluar ke `invoice_no`, sehingga jejak audit (siapa mengurangi stok apa dan kenapa) tetap utuh.
- `sales_orders.status` mendukung alur DRAFT → ISSUED → VOID, penting untuk pembatalan faktur tanpa menghapus data (audit trail).
- Tabel `tax_configs` menghindari hardcode tarif PPN di kode; backend membaca tarif aktif berdasarkan `effective_date`.

---

## 4. Desain REST API (Backend)

### 4.1 Autentikasi
| Method | Endpoint | Deskripsi |
| :--- | :--- | :--- |
| POST | `/api/v1/auth/login` | Login, kembalikan JWT access & refresh token |
| POST | `/api/v1/auth/refresh` | Refresh access token |
| POST | `/api/v1/auth/logout` | Invalidasi refresh token |

### 4.2 Master Data
| Method | Endpoint | Deskripsi |
| :--- | :--- | :--- |
| GET/POST | `/api/v1/partners` | List (dengan filter `type`) / Create partner |
| GET/PUT/DELETE | `/api/v1/partners/:id` | Detail / Update / Soft delete |
| GET/POST | `/api/v1/products` | List (search, pagination) / Create produk |
| GET/PUT/DELETE | `/api/v1/products/:sku` | Detail / Update / Soft delete |

### 4.3 Inventory
| Method | Endpoint | Deskripsi |
| :--- | :--- | :--- |
| POST | `/api/v1/inventory/stock-in` | Input barang masuk → trigger kalkulasi HPP Moving Average |
| POST | `/api/v1/inventory/adjustment` | Penyesuaian stok manual (opname) |
| GET | `/api/v1/inventory/ledger/:sku` | Riwayat mutasi stok per produk |

### 4.4 Sales / Faktur
| Method | Endpoint | Deskripsi |
| :--- | :--- | :--- |
| POST | `/api/v1/sales` | Buat faktur baru (hitung DPP/PPN, kurangi stok, catat HPP snapshot) — **dalam 1 DB transaction** |
| GET | `/api/v1/sales` | List faktur (filter tanggal, partner, status) |
| GET | `/api/v1/sales/:invoiceNo` | Detail faktur lengkap dengan items |
| POST | `/api/v1/sales/:invoiceNo/void` | Batalkan faktur (kembalikan stok) |
| GET | `/api/v1/sales/:invoiceNo/print` | Data siap-cetak untuk halaman print Vue |

### 4.5 Contoh Response Terstruktur (Buat Faktur)
```json
POST /api/v1/sales
{
  "partner_id": "uuid-partner",
  "tax_include": false,
  "items": [
    { "product_sku": "SKU001", "qty": 5, "unit_price": 100000 }
  ]
}

Response 201:
{
  "invoice_no": "INV-2026-00123",
  "total_dpp": 500000,
  "total_ppn": 55000,
  "grand_total": 555000,
  "items": [ ... ]
}
```

---

## 5. Logika Inti Backend (Go)

### 5.1 Kalkulasi HPP — Moving Average (dalam DB Transaction)

Proses stock-in HARUS berjalan dalam satu transaksi SQL dengan row-level lock (`SELECT ... FOR UPDATE`) untuk mencegah *race condition* saat dua input stok terjadi bersamaan pada produk yang sama.

```go
func (s *inventoryService) StockIn(ctx context.Context, sku string, qty int, unitCost decimal.Decimal) error {
    return s.db.Transaction(func(tx *gorm.DB) error {
        var product domain.Product
        // Lock baris produk agar tidak terjadi race condition antar transaksi paralel
        if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
            Where("sku = ?", sku).First(&product).Error; err != nil {
            return err
        }

        oldStock := decimal.NewFromInt(int64(product.CurrentStock))
        newQty := decimal.NewFromInt(int64(qty))
        oldHPP := product.CurrentHPP

        // ((StokLama * HPPLama) + (StokMasuk * HargaBaru)) / (StokLama + StokMasuk)
        totalOld := oldStock.Mul(oldHPP)
        totalNew := newQty.Mul(unitCost)
        totalStock := oldStock.Add(newQty)

        newHPP := oldHPP
        if !totalStock.IsZero() {
            newHPP = totalOld.Add(totalNew).Div(totalStock)
        }

        product.CurrentStock += qty
        product.CurrentHPP = newHPP
        if err := tx.Save(&product).Error; err != nil {
            return err
        }

        trx := domain.InventoryTransaction{
            ProductSKU: sku, TrxType: "IN", Qty: qty, UnitCost: unitCost,
        }
        return tx.Create(&trx).Error
    })
}
```

### 5.2 Kalkulasi Pajak (Standalone, tanpa hit API DJP)

```go
func CalculateTax(price decimal.Decimal, qty int, rate decimal.Decimal, taxInclude bool) (dpp, ppn decimal.Decimal) {
    q := decimal.NewFromInt(int64(qty))
    if taxInclude {
        // DPP = (Harga / (1 + Tarif)) x Qty
        divisor := decimal.NewFromInt(1).Add(rate)
        dppPerUnit := price.Div(divisor)
        dpp = dppPerUnit.Mul(q).Round(2)
        ppn = price.Mul(q).Sub(dpp).Round(2)
    } else {
        // DPP = Harga x Qty ; PPN = DPP x Tarif
        dpp = price.Mul(q).Round(2)
        ppn = dpp.Mul(rate).Round(2)
    }
    return dpp, ppn
}
```

### 5.3 Pembuatan Faktur Penjualan (Sales Service)

Alur dalam satu DB transaction:
1. Lock & validasi stok cukup untuk setiap item.
2. Ambil tarif pajak aktif dari `tax_configs`.
3. Hitung DPP/PPN per item, snapshot `unit_hpp` dari `products.current_hpp` saat ini.
4. Insert `sales_orders` + `sales_items`.
5. Insert `inventory_transactions` tipe `OUT` dan kurangi `products.current_stock`.
6. Generate `invoice_no` berurut (format `INV-{YYYY}-{sequence}`) menggunakan PostgreSQL sequence per tahun agar aman dari duplikasi saat concurrent request.

### 5.4 Generate Nomor Faktur Aman (Concurrent-Safe)

```sql
CREATE SEQUENCE IF NOT EXISTS invoice_seq_2026 START 1;
-- backend ambil nextval() dalam transaction yang sama saat create faktur
```

---

## 6. Frontend: Halaman Cetak A4 (Vue.js)

- Route khusus `PrintLayout.vue` tanpa sidebar/navbar, didaftarkan terpisah di `router/index.js` dengan `meta: { layout: 'print' }`.
- Komponen `PrintInvoiceView.vue` fetch data dari `GET /api/v1/sales/:invoiceNo/print`, lalu render tabel item + kop surat perusahaan.
- CSS print sesuai draft awal (`@page { size: A4; margin: 15mm }`, `page-break-inside: avoid` pada `<tr>`) dipasang lewat `<style scoped>` khusus di komponen ini, atau file global `print.css` yang di-import kondisional.
- Tombol "Cetak" memanggil `window.print()`, dengan `.no-print` menyembunyikan tombol navigasi saat proses cetak berjalan.

---

## 7. Non-Fungsional & Keamanan

| Aspek | Implementasi |
| :--- | :--- |
| Validasi input | `validator` di Go (struct tag) + `vee-validate`/`yup` di Vue (client-side) |
| Otorisasi | Middleware role-based (admin vs staff) di Gin |
| Audit trail | Kolom `created_by`, `created_at` di setiap tabel transaksi |
| Konsistensi data | Semua operasi tulis stok & faktur dibungkus DB transaction + row lock |
| Precision uang | `DECIMAL` di Postgres, `shopspring/decimal` di Go, hindari `float64` |
| Rate limiting | Middleware Gin (`gin-contrib/limiter` atau custom) untuk endpoint publik |
| CORS | `gin-contrib/cors` dikonfigurasi hanya untuk origin frontend |
| Environment secrets | `.env` (JWT secret, DB credentials) tidak masuk repo, gunakan `.env.example` |

---

## 8. Rencana Deployment

**docker-compose.yml (ringkasan struktur):**
```yaml
services:
  db:
    image: postgres:16
    volumes: [pgdata:/var/lib/postgresql/data]
    environment: [POSTGRES_DB, POSTGRES_USER, POSTGRES_PASSWORD]
  api:
    build: ./backend
    depends_on: [db]
    environment: [DB_DSN, JWT_SECRET]
  web:
    build: ./frontend
    depends_on: [api]
  nginx:
    image: nginx:alpine
    ports: ["80:80", "443:443"]
    depends_on: [web, api]
volumes:
  pgdata:
```

- **Migrasi:** jalankan `golang-migrate` sebagai *init container* atau job terpisah sebelum `api` start.
- **Backup:** cron job `pg_dump` harian ke storage terpisah (krusial karena data keuangan).
- **Monitoring dasar:** logging terstruktur (zap → stdout) + healthcheck endpoint `/healthz` di Gin.

---

## 9. Rencana Pengembangan Bertahap (Sprint Plan)

| Fase | Durasi (estimasi) | Cakupan |
| :--- | :--- | :--- |
| **Fase 1 — Fondasi** | 1–2 minggu | Setup repo, Docker Compose, migrasi DB, autentikasi JWT, CRUD `partners` & `products` |
| **Fase 2 — Inventory** | 1–2 minggu | Stock-in, kalkulasi HPP Moving Average, buku besar stok, halaman opname/adjustment |
| **Fase 3 — Penjualan & Pajak** | 2 minggu | Kalkulasi pajak include/exclude, pembuatan faktur (transaksional), pengurangan stok otomatis, void faktur |
| **Fase 4 — Cetak & Laporan** | 1 minggu | Layout cetak A4, laporan penjualan/stok sederhana (rekap harian/bulanan) |
| **Fase 5 — Hardening & Deploy** | 1 minggu | Testing (unit + integration), rate limiting, backup strategy, deployment produksi |

---

## 10. Strategi Testing

- **Backend:** unit test untuk `tax_service` dan `inventory_service` (kasus HPP dan pajak wajib punya test table-driven dengan angka presisi 2 desimal), integration test untuk endpoint `sales` menggunakan test database terpisah.
- **Frontend:** component test (Vitest + Vue Test Utils) untuk form validasi, e2e test (Playwright/Cypress) untuk alur: login → buat faktur → cetak.
- **Data-critical path:** setiap perubahan pada `tax_service.go` dan `inventory_service.go` wajib disertai test baru sebelum merge, karena keduanya menyentuh langsung angka keuangan.

---

### Ringkasan Perubahan dari Draft Awal
1. Backend ORM/query direkomendasikan `gorm` + `shopspring/decimal` untuk presisi uang (bukan float).
2. Tabel `tax_configs` ditambahkan agar tarif PPN tidak hardcode.
3. Kolom `status`, `reference_no`, `created_by` ditambahkan untuk audit trail.
4. Ditambahkan strategi row-locking untuk mencegah race condition pada stock-in dan pembuatan faktur secara concurrent.
5. Ditambahkan rencana testing dan sprint plan bertahap agar pengembangan terukur.
