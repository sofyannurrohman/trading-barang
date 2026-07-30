package controllers

import (
	"backend/db"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// --- Product Controllers ---

func GetProducts(c *gin.Context) {
	var products []models.Product
	if err := db.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}
	c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := db.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&product)
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

// --- Partner Controllers ---

func GetPartners(c *gin.Context) {
	var partners []models.Partner
	partnerType := c.Query("type") // Optional filter by type (client/supplier)

	query := db.DB
	if partnerType != "" {
		query = query.Where("type = ?", partnerType)
	}

	if err := query.Find(&partners).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch partners"})
		return
	}
	c.JSON(http.StatusOK, partners)
}

func CreatePartner(c *gin.Context) {
	var partner models.Partner
	if err := c.ShouldBindJSON(&partner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Create(&partner).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create partner"})
		return
	}
	c.JSON(http.StatusCreated, partner)
}

func UpdatePartner(c *gin.Context) {
	id := c.Param("id")
	var partner models.Partner
	if err := db.DB.First(&partner, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Partner not found"})
		return
	}

	if err := c.ShouldBindJSON(&partner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&partner)
	c.JSON(http.StatusOK, partner)
}

func DeletePartner(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.Partner{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete partner"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Partner deleted successfully"})
}
