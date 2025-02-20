package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Ambil URL dari environment variable
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("Database URL is not set in environment variables")
	}

	// Connect ke database PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	log.Println("Database connected!")
}
