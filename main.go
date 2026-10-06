package main

import (
	"product-catalog/controllers"
	"product-catalog/database"
	"product-catalog/middleware"
	"product-catalog/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Hubungkan ke Database PostgreSQL
	database.ConnectDB()

	// Migrasi otomatis tabel User dan Product
	database.DB.AutoMigrate(&models.User{}, &models.Product{})

	// Jalankan Seeder otomatis (seperti db/seeds.rb di Rails)
	database.SeedData()

	// Inisialisasi Gin Router
	r := gin.Default()

	// Mengaktifkan CORS (Cross-Origin Resource Sharing)
	r.Use(cors.Default())

	// --- ROUTE PUBLIK ---
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong! Server backend siap."})
	})

	// Grup API Publik (Auth & Get Product)
	api := r.Group("/api")
	{
		// Endpoint Autentikasi (sekarang jadi /api/login dan /api/register)
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		// Route Produk Publik
		api.GET("/products", controllers.GetProducts)
		api.GET("/products/:id", controllers.GetProductByID)
	}

	// --- ROUTE PROTECTED (Butuh Token JWT di bawah /api) ---
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
