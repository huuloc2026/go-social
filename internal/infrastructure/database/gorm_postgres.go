package database

import (
	"fmt"
	"log"

	"github.com/huuloc2026/go-social/config"
	"github.com/huuloc2026/go-social/internal/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
	// Build the Data Source Name (DSN)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	// Open a connection to the PostgreSQL database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Perform AutoMigrate to create the schema for the models
	models := []interface{}{
		&entities.User{},
		&entities.Post{},
		&entities.Like{},
		// Add other models as needed
	}

	// Automatically migrate the schema
	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return nil, fmt.Errorf("failed to migrate model %T: %w", model, err)
		}
	}

	// Log successful database connection
	log.Println("Database connection successfully established")

	// Return the database connection object
	return db, nil
}
