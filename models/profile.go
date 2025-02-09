package models

import (
	"time"

	"gorm.io/gorm"
)

// Profile represents the profiles table
type Profile struct {
	ID                 string    `gorm:"primaryKey" json:"id"`
	Email              string    `json:"email"`
	Name               string    `json:"name"`
	CompanyName        *string   `json:"company_name,omitempty"`
	ProfileImage       *string   `json:"profile_image,omitempty"`
	Phone              *string   `json:"phone,omitempty"`
	Role               *string   `json:"role,omitempty"`
	CMSEnabled         *bool     `json:"cms_enabled,omitempty"`
	GoogleAnalyticsKey *string   `json:"google_analytics_key,omitempty"`
	AhrefsKey          *string   `json:"ahrefs_key,omitempty"`
	SemrushKey         *string   `json:"semrush_key,omitempty"`
	OrganizationID     *string   `json:"organization_id,omitempty"`
	CreatedAt          time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relationships
	Organization *Organization        `gorm:"foreignKey:OrganizationID"`
	Memberships  []OrganizationMember `gorm:"foreignKey:UserID"`
}

// MigrateDB creates the table
func MigrateProfile(db *gorm.DB) {
	db.AutoMigrate(&Profile{})
}
