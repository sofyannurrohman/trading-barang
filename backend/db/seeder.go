package db

import (
	"backend/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RunAutoSeed() {
	var count int64
	DB.Model(&models.User{}).Count(&count)
	if count > 0 {
		fmt.Println("Database already contains users. Skipping auto-seed.")
		return
	}

	fmt.Println("Running Auto-Seed...")

	seedUser(DB)
	seedCompanyProfile(DB)
	partners := seedPartners(DB)
	products := seedProducts(DB)
	seedInventory(DB, products)
	seedSales(DB, partners, products)

	fmt.Println("Auto-Seed Complete! Admin credentials: admin / password")
}

func seedUser(d *gorm.DB) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	user := models.User{
		Username: "admin",
		Email:    "admin@tradingbarang.com",
		Password: string(hashedPassword),
		Role:     "admin",
	}
	d.Create(&user)
}

func seedCompanyProfile(d *gorm.DB) {
	profile := models.CompanyProfile{
		Name:          "UD DUO SRIKANDI",
		Address:       "Jl. Perdagangan Raya No. 88, Jawa Timur",
		NPWP:          "31.456.789.0-604.000",
		Phone:         "+62 812-3456-7890",
		Email:         "kontak@duosrikandi.com",
		BankInfo:      "BCA: 8870-123-456 a/n UD DUO SRIKANDI\nBRI: 0123-01-000456-50-1 a/n UD DUO SRIKANDI",
		InvoiceFooter: "Barang yang sudah dibeli tidak dapat ditukar/dikembalikan kecuali ada perjanjian tertulis sebelumnya. Pembayaran dianggap lunas apabila telah masuk ke rekening resmi kami.",
		PPNRate:       11.0,
	}
	d.Create(&profile)
}

func seedPartners(d *gorm.DB) map[string]models.Partner {
	supplier := models.Partner{
		Type:    "supplier",
		Name:    "CV Material Alam",
		Address: "Jl. Raya Bogor KM 30",
		NPWP:    "98.765.432.1-123.000",
	}
	d.Create(&supplier)

	client1 := models.Partner{
		Type:    "client",
		Name:    "PT Maju Konstruksi",
		Address: "Kawasan Industri Pulo Gadung",
		NPWP:    "11.222.333.4-555.000",
	}
	d.Create(&client1)

	client2 := models.Partner{
		Type:    "client",
		Name:    "Bpk. Budi Santoso",
		Address: "Perumahan Indah No 4",
		NIK:     "3201234567890001",
	}
	d.Create(&client2)

	return map[string]models.Partner{
		"supplier": supplier,
		"client1":  client1,
		"client2":  client2,
	}
}

func seedProducts(d *gorm.DB) map[string]models.Product {
	p1 := models.Product{
		SKU:           "SMN-50",
		Name:          "Semen Portland 50kg",
		Category:      "Bahan Bangunan",
		StandardPrice: 65000,
	}
	d.Create(&p1)

	p2 := models.Product{
		SKU:           "BS-12",
		Name:          "Besi Beton 12mm",
		Category:      "Bahan Bangunan",
		StandardPrice: 85000,
	}
	d.Create(&p2)

	return map[string]models.Product{
		"semen": p1,
		"besi":  p2,
	}
}

func seedInventory(d *gorm.DB, products map[string]models.Product) {
	semen := products["semen"]
	semenTx := models.InventoryTransaction{
		ProductID: semen.ID,
		Type:      "INBOUND",
		Quantity:  100,
		UnitPrice: 50000,
		Reference: "PO-001 (CV Material Alam)",
	}
	d.Create(&semenTx)
	semen.CurrentStock = 100
	semen.AverageHPP = 50000
	d.Save(&semen)

	besi := products["besi"]
	besiTx := models.InventoryTransaction{
		ProductID: besi.ID,
		Type:      "INBOUND",
		Quantity:  200,
		UnitPrice: 70000,
		Reference: "PO-002 (CV Material Alam)",
	}
	d.Create(&besiTx)
	besi.CurrentStock = 200
	besi.AverageHPP = 70000
	d.Save(&besi)

	products["semen"] = semen
	products["besi"] = besi
}

func seedSales(d *gorm.DB, partners map[string]models.Partner, products map[string]models.Product) {
	semen := products["semen"]
	besi := products["besi"]
	client1 := partners["client1"]

	inv1 := models.Invoice{
		InvoiceNumber: "INV-2026-0001",
		PartnerID:     client1.ID,
		TotalDPP:      (10 * semen.StandardPrice) + (20 * besi.StandardPrice),
		Status:        "PAID",
	}
	inv1.TotalPPN = inv1.TotalDPP * 0.11
	inv1.GrandTotal = inv1.TotalDPP + inv1.TotalPPN

	d.Create(&inv1)

	item1 := models.InvoiceItem{
		InvoiceID: inv1.ID,
		ProductID: semen.ID,
		Quantity:  10,
		UnitPrice: semen.StandardPrice,
		Subtotal:  10 * semen.StandardPrice,
		HPPAtSale: semen.AverageHPP,
	}
	d.Create(&item1)

	item2 := models.InvoiceItem{
		InvoiceID: inv1.ID,
		ProductID: besi.ID,
		Quantity:  20,
		UnitPrice: besi.StandardPrice,
		Subtotal:  20 * besi.StandardPrice,
		HPPAtSale: besi.AverageHPP,
	}
	d.Create(&item2)

	d.Create(&models.InventoryTransaction{
		ProductID: semen.ID,
		Type:      "SALE",
		Quantity:  -10,
		Reference: inv1.InvoiceNumber,
	})
	semen.CurrentStock -= 10
	d.Save(&semen)

	d.Create(&models.InventoryTransaction{
		ProductID: besi.ID,
		Type:      "SALE",
		Quantity:  -20,
		Reference: inv1.InvoiceNumber,
	})
	besi.CurrentStock -= 20
	d.Save(&besi)
}
