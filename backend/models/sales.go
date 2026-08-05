package models

import "gorm.io/gorm"

type Invoice struct {
	gorm.Model
	InvoiceNumber string        `gorm:"uniqueIndex;not null" json:"invoice_number"`
	PartnerID     uint          `gorm:"not null" json:"partner_id"`
	TotalDPP      float64       `json:"total_dpp"`
	TotalPPN      float64       `json:"total_ppn"`
	GrandTotal    float64       `json:"grand_total"`
	ShippingCost  float64       `gorm:"default:0" json:"shipping_cost"`
	Discount      float64       `gorm:"default:0" json:"discount"`
	Status        string        `gorm:"default:'UNPAID'" json:"status"`
	IsTaxable     bool          `gorm:"default:true" json:"is_taxable"`
	Items         []InvoiceItem `gorm:"foreignKey:InvoiceID" json:"items"`
	Partner       Partner       `gorm:"foreignKey:PartnerID" json:"partner"`
}

type InvoiceItem struct {
	gorm.Model
	InvoiceID   uint    `json:"invoice_id"`
	ProductID   uint    `json:"product_id"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"` // DPP per unit
	Subtotal    float64 `json:"subtotal"`
	HPPAtSale   float64 `json:"hpp_at_sale"`
	Product     Product `gorm:"foreignKey:ProductID" json:"product"`
}
