package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Sesuaikan "postgres" (user), "secretpassword" (password), dan "product_db" (nama database)
	// dengan yang Anda buat saat instalasi PostgreSQL.
	dsn := "host=localhost user=postgres password=postgres dbname=product_db port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database: ", err)
	}

	fmt.Println("Berhasil terhubung ke database PostgreSQL!")
}
