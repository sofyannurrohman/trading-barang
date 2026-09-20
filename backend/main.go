package main

import (
	"backend/controllers"
	"backend/db"
	"backend/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set Gin mode from environment (use GIN_MODE=release in production)
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize Database
	db.InitDB()

	if os.Getenv("AUTO_SEED") == "true" {
		db.RunAutoSeed()
	}

	r := gin.Default()
	r.Use(middleware.CORS())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	auth := r.Group("/auth")
	{
		// NOTE: /auth/register has been intentionally removed.
		// User creation is managed by admins via POST /api/settings/users (requires authentication).
		auth.POST("/login", controllers.Login)
	}

	api := r.Group("/api")
	api.Use(middleware.AuthRequired())
	{
		// Dashboard stats (real data)
		api.GET("/dashboard", controllers.Dashboard)

		// Master Data - Products
		api.GET("/products", controllers.GetProducts)
		api.POST("/products", controllers.CreateProduct)
		api.PUT("/products/:id", controllers.UpdateProduct)
		api.DELETE("/products/:id", controllers.DeleteProduct)
		api.POST("/products/:id/image", controllers.UploadProductImage)

		// Master Data - Partners
		api.GET("/partners", controllers.GetPartners)
		api.POST("/partners", controllers.CreatePartner)
		api.PUT("/partners/:id", controllers.UpdatePartner)
		api.DELETE("/partners/:id", controllers.DeletePartner)

		// Inventory
		api.POST("/inventory/inbound", controllers.CreateInbound)
		api.POST("/inventory/adjustment", controllers.CreateAdjustment)
		api.GET("/inventory/stock-card/:product_id", controllers.GetStockCard)

		// Sales
		api.POST("/sales/invoice", controllers.CreateInvoice)
		api.GET("/sales/invoices", controllers.GetInvoices)
		api.GET("/sales/invoices/:id", controllers.GetInvoiceDetails)
		api.PUT("/sales/invoices/:id", controllers.UpdateInvoice)
		api.DELETE("/sales/invoices/:id", controllers.DeleteInvoice)
		api.POST("/sales/invoices/:id/restore", controllers.RestoreInvoice)

		// Reports
		api.GET("/reports/profit", controllers.GetProfitReport)
		api.GET("/reports/tax", controllers.GetTaxReport)

		// Settings
		api.GET("/settings/company", controllers.GetCompanyProfile)
		api.PUT("/settings/company", controllers.UpdateCompanyProfile)

		// Users (admin only)
		api.GET("/settings/users", middleware.RoleRequired("admin"), controllers.GetUsers)
		api.POST("/settings/users", middleware.RoleRequired("admin"), controllers.CreateUser)
		api.PUT("/settings/users/:id", middleware.RoleRequired("admin"), controllers.UpdateUser)
		api.DELETE("/settings/users/:id", middleware.RoleRequired("admin"), controllers.DeleteUser)
		api.POST("/settings/users/:id/restore", middleware.RoleRequired("admin"), controllers.RestoreUser)
	}

	// Serve uploaded static files
	r.Static("/uploads", "./uploads")

	r.Run(":8080")
}
