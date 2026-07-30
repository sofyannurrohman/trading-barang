package controllers

import (
	"backend/db"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfitReportItem struct {
	InvoiceNumber string  `json:"invoice_number"`
	Date          string  `json:"date"`
	Revenue       float64 `json:"revenue"`
	TotalCost     float64 `json:"total_cost"`
	GrossProfit   float64 `json:"gross_profit"`
	MarginPercent float64 `json:"margin_percent"`
}

func GetProfitReport(c *gin.Context) {
	var invoices []models.Invoice
	if err := db.DB.Preload("Items").Order("created_at desc").Find(&invoices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch invoices"})
		return
	}

	var report []ProfitReportItem

	for _, inv := range invoices {
		var totalCost float64 = 0
		for _, item := range inv.Items {
			// Cost = Quantity * HPP at the time of sale
			totalCost += float64(item.Quantity) * item.HPPAtSale
		}

		revenue := inv.TotalDPP
		grossProfit := revenue - totalCost
		var margin float64 = 0
		if revenue > 0 {
			margin = (grossProfit / revenue) * 100
		}

		report = append(report, ProfitReportItem{
			InvoiceNumber: inv.InvoiceNumber,
			Date:          inv.CreatedAt.Format("2006-01-02 15:04:05"),
			Revenue:       revenue,
			TotalCost:     totalCost,
			GrossProfit:   grossProfit,
			MarginPercent: margin,
		})
	}

	c.JSON(http.StatusOK, report)
}

type TaxReportItem struct {
	InvoiceNumber string  `json:"invoice_number"`
	Date          string  `json:"date"`
	ClientName    string  `json:"client_name"`
	NPWPNIK       string  `json:"npwp_nik"`
	DPP           float64 `json:"dpp"`
	PPN           float64 `json:"ppn"`
}

func GetTaxReport(c *gin.Context) {
	var invoices []models.Invoice
	if err := db.DB.Preload("Partner").Order("created_at desc").Find(&invoices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch invoices"})
		return
	}

	var report []TaxReportItem

	for _, inv := range invoices {
		identifier := inv.Partner.NPWP
		if identifier == "" {
			identifier = inv.Partner.NIK + " (NIK)"
		}

		report = append(report, TaxReportItem{
			InvoiceNumber: inv.InvoiceNumber,
			Date:          inv.CreatedAt.Format("2006-01-02 15:04:05"),
			ClientName:    inv.Partner.Name,
			NPWPNIK:       identifier,
			DPP:           inv.TotalDPP,
			PPN:           inv.TotalPPN,
		})
	}

	c.JSON(http.StatusOK, report)
}
