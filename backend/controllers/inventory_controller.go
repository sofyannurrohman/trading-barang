package controllers

import (
	"backend/db"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InboundInput struct {
	ProductID uint    `json:"product_id" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required,gt=0"`
	UnitPrice float64 `json:"unit_price" binding:"required,gt=0"`
	Reference string  `json:"reference"`
}

func CreateInbound(c *gin.Context) {
	var input InboundInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := db.DB.Begin()

	// Find the product
	var product models.Product
	if err := tx.First(&product, input.ProductID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Calculate New HPP
	// New HPP = ((Current Stock * Current HPP) + (Inbound Qty * Inbound Price)) / (Current Stock + Inbound Qty)
	currentTotalValue := float64(product.CurrentStock) * product.AverageHPP
	inboundTotalValue := float64(input.Quantity) * input.UnitPrice
	newTotalStock := product.CurrentStock + input.Quantity
	
	newAverageHPP := (currentTotalValue + inboundTotalValue) / float64(newTotalStock)

	// Update product
	product.CurrentStock = newTotalStock
	product.AverageHPP = newAverageHPP

	if err := tx.Save(&product).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product stock and HPP"})
		return
	}

	// Record transaction
	transaction := models.InventoryTransaction{
		ProductID: input.ProductID,
		Type:      "INBOUND",
		Quantity:  input.Quantity,
		UnitPrice: input.UnitPrice,
		Reference: input.Reference,
	}

	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record transaction"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusCreated, gin.H{
		"message": "Inbound recorded successfully",
		"product": product,
		"transaction": transaction,
	})
}

type AdjustmentInput struct {
	ProductID uint   `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"` // Can be negative
	Reference string `json:"reference"`
}

func CreateAdjustment(c *gin.Context) {
	var input AdjustmentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := db.DB.Begin()

	var product models.Product
	if err := tx.First(&product, input.ProductID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Ensure stock doesn't go negative if it's a deduction
	newTotalStock := product.CurrentStock + input.Quantity
	if newTotalStock < 0 {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Stock cannot be negative"})
		return
	}

	product.CurrentStock = newTotalStock
	if err := tx.Save(&product).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product stock"})
		return
	}

	transaction := models.InventoryTransaction{
		ProductID: input.ProductID,
		Type:      "ADJUSTMENT",
		Quantity:  input.Quantity,
		UnitPrice: product.AverageHPP, // Record current HPP at time of adjustment for reporting
		Reference: input.Reference,
	}

	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record transaction"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusCreated, gin.H{
		"message": "Adjustment recorded successfully",
		"product": product,
		"transaction": transaction,
	})
}

func GetStockCard(c *gin.Context) {
	productID := c.Param("product_id")
	
	var transactions []models.InventoryTransaction
	if err := db.DB.Where("product_id = ?", productID).Order("created_at asc").Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch stock card"})
		return
	}
	
	c.JSON(http.StatusOK, transactions)
}
