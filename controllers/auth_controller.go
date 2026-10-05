package controllers

import (
	"net/http"
	"product-catalog/database"
	"product-catalog/middleware"
	"product-catalog/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Register - Mendaftarkan user baru
func Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password sebelum disimpan ke database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengenkripsi password"})
		return
	}
	input.Password = string(hashedPassword)

	// Simpan ke database
	if err := database.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username sudah terdaftar atau tidak valid"})
		return
	}

	// Sembunyikan password pada response json
	input.Password = ""

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registrasi berhasil",
		"data":    input,
	})
}

// Login - Autentikasi user dan menghasilkan Token JWT
func Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	// Cari user berdasarkan username
	if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username atau password salah"})
		return
	}

	// Verifikasi password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username atau password salah"})
		return
	}

	// Buat klaim JWT (masa aktif token misal 2 jam)
	expirationTime := time.Now().Add(2 * time.Hour)
	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(middleware.JwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghasilkan token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login berhasil",
		"token":   tokenString,
	})
}
