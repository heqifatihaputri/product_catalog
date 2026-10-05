package main

import (
	"product-catalog/controllers"
	"product-catalog/database"
	"product-catalog/middleware"
	"product-catalog/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Hubungkan ke Database PostgreSQL
	database.ConnectDB()

	// 2. Migrasi otomatis tabel User dan Product
	database.DB.AutoMigrate(&models.User{}, &models.Product{})

	// 3. Inisialisasi Gin Router
	r := gin.Default()

	// --- ROUTE PUBLIK ---
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong! Server backend siap."})
	})

	// Endpoint Autentikasi
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Route Produk Publik (Contoh: Siapa saja boleh melihat daftar produk)
	r.GET("/products", controllers.GetProducts)
	r.GET("/products/:id", controllers.GetProductByID)

	// --- ROUTE PROTECTED (Butuh Token JWT) ---
	authRoutes := r.Group("/api")
	authRoutes.Use(middleware.AuthMiddleware())
	{
		authRoutes.POST("/products", controllers.CreateProduct)
		authRoutes.PUT("/products/:id", controllers.UpdateProduct)
		authRoutes.DELETE("/products/:id", controllers.DeleteProduct)
	}

	// Jalankan server di port 8080
	r.Run(":8080")
}
