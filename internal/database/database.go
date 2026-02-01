package database

import (
	"fmt"
	"log"
	"time"

	"github.com/terryluciano/templ-test/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	log.Println("Connecting to Database...")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", config.Config.DATABASE_HOST, config.Config.DATABASE_USER, config.Config.DATABASE_PASSWORD, config.Config.DATABASE_NAME, config.Config.DATABASE_PORT, config.Config.DATABASE_SSL_MODE, config.Config.DATABASE_TIMEZONE)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to Connect Database: %v", err)
	}

	psqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get database instance: %v", err)
	}

	psqlDB.SetMaxIdleConns(10)
	psqlDB.SetMaxOpenConns(100)
	psqlDB.SetConnMaxLifetime(time.Hour)

	DB = db

	log.Println("Database Connected")
}
