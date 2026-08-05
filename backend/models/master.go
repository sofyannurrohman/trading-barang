package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	SKU           string  `gorm:"uniqueIndex;not null" json:"sku"`
	Name          string  `gorm:"not null" json:"name"`
	Category      string  `json:"category"`
	StandardPrice float64 `gorm:"default:0" json:"standard_price"` // Harga Jual Standar
	AverageHPP    float64 `gorm:"default:0" json:"average_hpp"`    // Harga Pokok Penjualan (Modal rata-rata)
	CurrentStock  int     `gorm:"default:0" json:"current_stock"`
	ImageURL      string  `json:"image_url"` // Path relatif, e.g. "/uploads/products/abc.jpg"
}

type Partner struct {
	gorm.Model
	Type    string `gorm:"not null" json:"type"` // "client" or "supplier"
	Name    string `gorm:"not null" json:"name"`
	Address string `json:"address"`
	NPWP    string `json:"npwp"`
	NIK     string `json:"nik"` // Untuk pembeli yang tidak punya NPWP
}
