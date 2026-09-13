package db

import (
	"backend/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslMode := os.Getenv("DB_SSLMODE")
	if sslMode == "" {
		sslMode = "disable" // safe default for Docker internal networks
	}

	var dsn string
	if host != "" && user != "" && password != "" && dbname != "" && port != "" {
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
			host, user, password, dbname, port, sslMode,
		)
	} else if os.Getenv("DATABASE_URL") != "" {
		dsn = os.Getenv("DATABASE_URL")
	} else {
		// Fallback for local development without Docker
		dsn = "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=UTC"
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
