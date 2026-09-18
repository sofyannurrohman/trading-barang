package controllers

import (
	"backend/db"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCompanyProfile(c *gin.Context) {
	var profile models.CompanyProfile
	
	// Fetch the first profile. If it doesn't exist, return an initial UD Duo Srikandi profile.
	result := db.DB.First(&profile)
	if result.Error != nil {
		profile = models.CompanyProfile{
			Name:          "UD DUO SRIKANDI",
			Address:       "Jl. Perdagangan Raya No. 88, Jawa Timur",
			NPWP:          "31.456.789.0-604.000",
			Phone:         "+62 812-3456-7890",
			Email:         "kontak@duosrikandi.com",
			BankInfo:      "BCA: 8870-123-456 a/n UD DUO SRIKANDI\nBRI: 0123-01-000456-50-1 a/n UD DUO SRIKANDI",
			InvoiceFooter: "Barang yang sudah dibeli tidak dapat ditukar/dikembalikan kecuali ada perjanjian tertulis sebelumnya. Pembayaran dianggap lunas apabila telah masuk ke rekening resmi kami.",
			PPNRate:       11.0,
		}
		// Auto-save initial profile so it persists in DB
		db.DB.Create(&profile)
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
	profile.Phone = input.Phone
	profile.Email = input.Email
	profile.BankInfo = input.BankInfo
	profile.InvoiceFooter = input.InvoiceFooter
	profile.PPNRate = input.PPNRate
	profile.LogoURL = input.LogoURL

	if err := db.DB.Save(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update company profile"})
		return
	}

	c.JSON(http.StatusOK, profile)
}
