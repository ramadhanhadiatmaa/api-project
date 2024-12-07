package models

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	// Load environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	// Construct the DSN (Data Source Name) for MySQL connection
	dsn := os.Getenv("DB_USER") + ":" +
		os.Getenv("DB_PASS") + "@tcp(" +
		os.Getenv("DB_HOST") + ":" +
		os.Getenv("DB_PORT") + ")/" +
		os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

	// Open the database connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Perform database migration
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("Failed to migrate the database: %v", err)
	}

	// Assign the connection to the global variable
	DB = db
	log.Println("Successfully connected to the database.")
}
