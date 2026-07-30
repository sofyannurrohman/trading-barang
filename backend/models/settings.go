package models

import "gorm.io/gorm"

type CompanyProfile struct {
	gorm.Model
	Name    string  `json:"name"`
	Address string  `json:"address"`
	NPWP    string  `json:"npwp"`
	LogoURL string  `json:"logo_url"`
	PPNRate float64 `gorm:"default:11.0" json:"ppn_rate"` // Persentase PPN, misal 11
}
