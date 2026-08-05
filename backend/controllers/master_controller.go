package controllers

import (
	"backend/db"
	"backend/models"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func UploadProductImage(c *gin.Context) {
	id := c.Param("id")

	// Cari produk
	var product models.Product
	if err := db.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	// Ambil file dari form
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File tidak ditemukan di request"})
		return
	}
	defer file.Close()

	// Validasi ukuran maks 2MB
	if header.Size > 2*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ukuran file melebihi 2MB"})
		return
	}

	// Validasi MIME type dari header file
	buf := make([]byte, 512)
	if _, err := file.Read(buf); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca file"})
		return
	}
	mimeType := http.DetectContentType(buf)
	allowed := map[string]string{
		"image/jpeg": ".jpg",
		"image/png":  ".png",
		"image/webp": ".webp",
	}
	ext, ok := allowed[mimeType]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format file tidak didukung. Gunakan JPG, PNG, atau WebP"})
		return
	}

	// Buat direktori uploads jika belum ada
	uploadDir := "./uploads/products"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat direktori upload"})
		return
	}

	// Generate nama file unik
	shortUUID := uuid.New().String()[:8]
	filename := fmt.Sprintf("%d_%s%s", product.ID, shortUUID, ext)
	filePath := fmt.Sprintf("%s/%s", uploadDir, filename)

	// Hapus foto lama jika ada
	if product.ImageURL != "" {
		oldPath := "." + product.ImageURL
		os.Remove(oldPath)
	}

	// Seek kembali ke awal karena sudah dibaca 512 byte untuk MIME detection
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses file"})
		return
	}

	// Tulis file ke disk
	out, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
		return
	}
	defer out.Close()

	if _, err := io.Copy(out, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menulis file"})
		return
	}

	// Simpan URL ke database
	imageURL := "/uploads/products/" + filename
	if err := db.DB.Model(&product).Update("image_url", imageURL).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data foto"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"image_url": imageURL, "message": "Foto produk berhasil diperbarui"})
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
