package models

import "gorm.io/gorm"

type CompanyProfile struct {
	gorm.Model
	Name          string  `json:"name"`
	Address       string  `json:"address"`
	NPWP          string  `json:"npwp"`
	Phone         string  `json:"phone"`
	Email         string  `json:"email"`
	BankInfo      string  `json:"bank_info"`
	InvoiceFooter string  `json:"invoice_footer"`
	LogoURL       string  `json:"logo_url"`
	PPNRate       float64 `gorm:"default:11.0" json:"ppn_rate"`
}
