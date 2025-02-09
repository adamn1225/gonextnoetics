package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"nextnoetics.com/backend/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	fmt.Println("✅ Successfully connected to PostgreSQL!")

	// ✅ Auto-migrate all models (ensures relationships are created)
	db.AutoMigrate(
		&models.Organization{},
		&models.OrganizationMember{},
		&models.Profile{},
		&models.Task{}, // ✅ Now includes tasks with relationships
	)

	DB = db
}
