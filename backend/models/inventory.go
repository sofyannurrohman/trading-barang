package models

import "gorm.io/gorm"

type InventoryTransaction struct {
	gorm.Model
	ProductID uint    `gorm:"not null" json:"product_id"`
	Type      string  `gorm:"not null" json:"type"`      // "INBOUND", "ADJUSTMENT", "SALE"
	Quantity  int     `gorm:"not null" json:"quantity"`  // Can be negative
	UnitPrice float64 `gorm:"default:0" json:"unit_price"` // Used for INBOUND (purchase price)
	Reference string  `json:"reference"`                 // e.g., "Supplier PO #123"
}
