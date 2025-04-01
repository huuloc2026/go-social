package database

import (
	"fmt"
	"log"

	// Import the correct config package
	"github.com/huuloc2026/go-social/config"
	"github.com/huuloc2026/go-social/internal/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
	// Build the Data Source Name (DSN)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	// Log the generated DSN for debugging purposes
	log.Println("DSN:", dsn)
	// Open a connection to the PostgreSQL database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Perform AutoMigrate to create the schema
	if err := db.AutoMigrate(&entities.User{}); err != nil {
		return nil, err
	}

	// Log successful database connection
	log.Println("Database Connection successfully established")

	return db, nil
}
