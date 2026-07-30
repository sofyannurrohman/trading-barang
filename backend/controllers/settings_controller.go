package controllers

import (
	"backend/db"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCompanyProfile(c *gin.Context) {
	var profile models.CompanyProfile
	
	// Fetch the first profile. If it doesn't exist, return an empty one or default.
	result := db.DB.First(&profile)
	if result.Error != nil {
		// Just return a default one without saving to avoid issues if table is empty
		profile = models.CompanyProfile{
			Name:    "PT TRADING BARANG",
			Address: "Jl. Contoh Alamat No. 123",
			NPWP:    "00.000.000.0-000.000",
			PPNRate: 11.0,
		}
	}

	c.JSON(http.StatusOK, profile)
}

func UpdateCompanyProfile(c *gin.Context) {
	var input models.CompanyProfile
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var profile models.CompanyProfile
	if err := db.DB.First(&profile).Error; err != nil {
		// If none exists, create it
		if err := db.DB.Create(&input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create company profile"})
			return
		}
		c.JSON(http.StatusOK, input)
		return
	}

	// Update existing
	profile.Name = input.Name
	profile.Address = input.Address
	profile.NPWP = input.NPWP
	profile.PPNRate = input.PPNRate
	profile.LogoURL = input.LogoURL

	if err := db.DB.Save(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update company profile"})
		return
	}

	c.JSON(http.StatusOK, profile)
}
