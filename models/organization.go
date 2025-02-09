package models

import (
	"time"

	"gorm.io/gorm"
)

// Organization represents the organizations table
type Organization struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	// Relationships
	Members  []OrganizationMember `gorm:"foreignKey:OrganizationID"`
	Profiles []Profile            `gorm:"foreignKey:OrganizationID"`
}

// MigrateDB creates the table
func MigrateOrganization(db *gorm.DB) {
	db.AutoMigrate(&Organization{})
}
