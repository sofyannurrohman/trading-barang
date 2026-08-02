package controllers

import (
	"backend/db"
	"backend/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InvoiceItemInput struct {
	ProductID uint    `json:"product_id" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required,gt=0"`
	UnitPrice float64 `json:"unit_price" binding:"required,gt=0"` // Harga Jual per unit
}

type CreateInvoiceInput struct {
	PartnerID    uint               `json:"partner_id" binding:"required"`
	Items        []InvoiceItemInput `json:"items" binding:"required,min=1"`
	ShippingCost float64            `json:"shipping_cost"`
	Discount     float64            `json:"discount"`
}

func CreateInvoice(c *gin.Context) {
	var input CreateInvoiceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := db.DB.Begin()

	// Generate Invoice Number (Simple generator)
	invoiceNumber := fmt.Sprintf("INV-%d", time.Now().Unix())

	var totalDPP float64 = 0

	invoice := models.Invoice{
		InvoiceNumber: invoiceNumber,
		PartnerID:     input.PartnerID,
		ShippingCost:  input.ShippingCost,
		Discount:      input.Discount,
		Status:        "UNPAID",
	}

	if err := tx.Create(&invoice).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create invoice"})
		return
	}

	for _, itemInput := range input.Items {
		var product models.Product
		if err := tx.First(&product, itemInput.ProductID).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Product %d not found", itemInput.ProductID)})
			return
		}

		if product.CurrentStock < itemInput.Quantity {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Insufficient stock for product %s", product.Name)})
			return
		}

		subtotal := float64(itemInput.Quantity) * itemInput.UnitPrice
		totalDPP += subtotal

		// Deduct stock
		product.CurrentStock -= itemInput.Quantity
		if err := tx.Save(&product).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update stock"})
			return
		}

		// Create Invoice Item
		invoiceItem := models.InvoiceItem{
			InvoiceID: invoice.ID,
			ProductID: product.ID,
			Quantity:  itemInput.Quantity,
			UnitPrice: itemInput.UnitPrice,
			Subtotal:  subtotal,
			HPPAtSale: product.AverageHPP,
		}

		if err := tx.Create(&invoiceItem).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save invoice item"})
			return
		}

		// Record inventory transaction
		transaction := models.InventoryTransaction{
			ProductID: product.ID,
			Type:      "SALE",
			Quantity:  -itemInput.Quantity,
			UnitPrice: itemInput.UnitPrice,
			Reference: invoiceNumber,
		}

		if err := tx.Create(&transaction).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record transaction"})
			return
		}
	}

	// Fetch Company Profile to get PPN Rate (Fallback to 11% if not set)
	var profile models.CompanyProfile
	ppnRate := 11.0
	if err := tx.First(&profile).Error; err == nil && profile.PPNRate > 0 {
		ppnRate = profile.PPNRate
	}

	totalPPN := totalDPP * (ppnRate / 100.0)
	grandTotal := totalDPP + totalPPN

	invoice.TotalDPP = totalDPP
	invoice.TotalPPN = totalPPN
	invoice.GrandTotal = grandTotal + input.ShippingCost - input.Discount

	if err := tx.Save(&invoice).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to finalize invoice"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusCreated, invoice)
}

func GetInvoices(c *gin.Context) {
	var invoices []models.Invoice
	if err := db.DB.Preload("Partner").Order("created_at desc").Find(&invoices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch invoices"})
		return
	}
	c.JSON(http.StatusOK, invoices)
}

func GetInvoiceDetails(c *gin.Context) {
	id := c.Param("id")
	var invoice models.Invoice
	if err := db.DB.Preload("Partner").Preload("Items").Preload("Items.Product").First(&invoice, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}
	c.JSON(http.StatusOK, invoice)
}

type UpdateInvoiceInput struct {
	Status       string  `json:"status"`
	ShippingCost float64 `json:"shipping_cost"`
	Discount     float64 `json:"discount"`
}

func UpdateInvoice(c *gin.Context) {
	id := c.Param("id")
	var invoice models.Invoice
	if err := db.DB.First(&invoice, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice tidak ditemukan"})
		return
	}

	var input UpdateInvoiceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Status != "" {
		invoice.Status = input.Status
	}
	invoice.ShippingCost = input.ShippingCost
	invoice.Discount = input.Discount
	// Recalculate GrandTotal: DPP + PPN + Ongkir - Diskon
	invoice.GrandTotal = invoice.TotalDPP + invoice.TotalPPN + invoice.ShippingCost - invoice.Discount

	if err := db.DB.Save(&invoice).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan perubahan"})
		return
	}
	c.JSON(http.StatusOK, invoice)
}

func DeleteInvoice(c *gin.Context) {
	id := c.Param("id")
	var invoice models.Invoice
	if err := db.DB.Preload("Items").Preload("Partner").First(&invoice, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice tidak ditemukan"})
		return
	}

	tx := db.DB.Begin()

	// Rollback stok untuk setiap item di invoice ini
	for _, item := range invoice.Items {
		if err := tx.Model(&models.Product{}).Where("id = ?", item.ProductID).
			UpdateColumn("current_stock", gorm.Expr("current_stock + ?", -item.Quantity)).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengembalikan stok produk"})
			return
		}
	}

	// Soft-delete invoice (GORM DeletedAt)
	if err := tx.Delete(&invoice).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus invoice"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"message": "Invoice berhasil dihapus dan stok produk telah dikembalikan"})
}
