package models

import "time"

type OrganizationMember struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	CMSToken         *string   `json:"cms_token,omitempty"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
	OrganizationID   string    `gorm:"not null" json:"organization_id"`
	OrganizationName *string   `json:"organization_name,omitempty"`
	Role             string    `json:"role"`
	UserID           string    `gorm:"not null" json:"user_id"`
	WebsiteURL       *string   `json:"website_url,omitempty"`

	// ✅ Define Foreign Key Relationships Correctly
	Organization Organization `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"organization"`
	Profile      Profile      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"profile"`
}
