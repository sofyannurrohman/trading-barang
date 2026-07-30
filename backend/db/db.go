package db

import (
	"backend/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=UTC" // default for local testing if not provided
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto Migrate
	err = DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Partner{},
		&models.CompanyProfile{},
		&models.InventoryTransaction{},
		&models.Invoice{},
		&models.InvoiceItem{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
