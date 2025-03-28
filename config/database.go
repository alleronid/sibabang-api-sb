package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func ConnectDB() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	name := os.Getenv("DB_NAME")
	cloudhost := os.Getenv("DB_HOST")
	cloudport := os.Getenv("DB_PORT")
	cloudusername := os.Getenv("DB_USERNAME")
	cloudpassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cloudusername, cloudpassword, cloudhost, cloudport, name)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
