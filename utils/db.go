package utils

import (
    "log"
    "employees-app/models"
    "os"
    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func InitDB() *gorm.DB {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Get database configuration from environment variables
    dbHost := os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbPort := os.Getenv("DB_PORT")

    // Membuat Data Source Name (DSN)
    dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable"
    
    // Menghubungkan ke database
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database, got error: %v", err)
    }

    // Menjalankan migrasi
    MigrateDB(db)

    return db
}

func MigrateDB(db *gorm.DB) {
    err := db.AutoMigrate(&models.Employee{})
    if err != nil {
        log.Fatalf("failed to migrate database, got error: %v", err)
    }
}
