package models

import (
	"time"

	"gorm.io/gorm"
)

// Task represents the tasks table
type Task struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	Title          string     `json:"title"`
	Description    string     `json:"description"`
	DueDate        *time.Time `json:"due_date,omitempty"` // Nullable datetime
	IsRequest      bool       `json:"is_request"`
	Status         string     `json:"status"`
	TaskType       string     `gorm:"default:'client'" json:"task_type"` // ✅ Default value set to "client"
	OrganizationID string     `gorm:"not null" json:"organization_id"`
	UserID         string     `gorm:"not null" json:"user_id"`
	CreatedAt      time.Time  `gorm:"autoCreateTime" json:"created_at"`

	// ✅ Relationships
	Profile      Profile      `gorm:"foreignKey:UserID;references:ID"`
	Organization Organization `gorm:"foreignKey:OrganizationID;references:ID"`
}

// MigrateDB creates the tasks table
func MigrateTask(db *gorm.DB) {
	db.AutoMigrate(&Task{})
}
