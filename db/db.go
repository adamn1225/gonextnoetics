package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"nextnoetics.com/backend/models"
)

// DB instance
var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=myuser password=Adam123! dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	DB = db
	fmt.Println("✅ Successfully connected to PostgreSQL!")

	// Run migrations
	models.MigrateDB(DB)
}
