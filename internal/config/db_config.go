package config

import (
	"log"

	"github.com/Amotekun/payment-system/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:codewithamotekun#@tcp(127.0.0.1:3306)/payment_system?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connec to database: %v", err)
	}
	log.Println("Databse connection established")

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
}
