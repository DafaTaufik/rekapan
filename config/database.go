package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
    
    host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

    // Check if the environment variables are set
    if host == "" || port == "" || user == "" || dbname == "" {
        log.Fatal("env is not complete, Check .env file")
    }
    
    // Create the connection string
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",host, port, user, password, dbname) 
    
    // Connect to the database
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

	DB = database
	fmt.Println("Database connected successfully")
}