package controllers

import (
	"net/http"
	"product-catalog/database"
	"product-catalog/models"

	"github.com/gin-gonic/gin"
)

// GetProducts - Mengambil semua data produk
func GetProducts(c *gin.Context) {
	var products []models.Product
	if err := database.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data produk"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})
}

// GetProductByID - Mengambil detail satu produk berdasarkan ID
func GetProductByID(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// CreateProduct - Menambahkan produk baru
func CreateProduct(c *gin.Context) {
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&input)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Produk berhasil ditambahkan",
		"data":    input,
	})
}

// UpdateProduct - Mengubah data produk berdasarkan ID
func UpdateProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	// Cek apakah produknya ada
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	// Tangkap data inputan baru
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update data di database
	database.DB.Model(&product).Updates(models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "Produk berhasil diperbarui",
		"data":    product,
	})
}

// DeleteProduct - Menghapus produk berdasarkan ID
func DeleteProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	database.DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil dihapus"})
}
