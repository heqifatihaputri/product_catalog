package database

import (
	"product-catalog/models"

	"golang.org/x/crypto/bcrypt"
)

// SeedData menjalankan seeding data awal jika diperlukan
func SeedData() {
	var existingAdmin models.User
	// Cek apakah user admin sudah ada
	if err := DB.Where("username = ?", "admin").First(&existingAdmin).Error; err != nil {
		// Jika belum ada, buat otomatis dengan password ter-hash bcrypt
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("passwordBaru123"), bcrypt.DefaultCost)
		newAdmin := models.User{
			Username: "admin",
			Password: string(hashedPassword),
		}
		DB.Create(&newAdmin)
	}
}
