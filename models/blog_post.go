package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

// BlogPost represents the blog_posts table
type BlogPost struct {
	ID                   uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Title                string         `json:"title"`
	Content              string         `json:"content"`
	ContentHTML          *string        `json:"content_html,omitempty"`
	Excerpt              *string        `json:"excerpt,omitempty"`
	FeaturedImage        *string        `json:"featured_image,omitempty"`
	PublishedAt          *time.Time     `json:"published_at,omitempty"`
	ScheduledPublishDate *time.Time     `json:"scheduled_publish_date,omitempty"`
	Status               string         `json:"status"`
	Template             string         `json:"template"`
	Slug                 string         `gorm:"uniqueIndex" json:"slug"`
	CustomFields         pq.StringArray `gorm:"type:text[]" json:"customFields"`
	UserID               *string        `json:"user_id,omitempty"`
	CreatedAt            time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt            time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}

// MigrateDB creates the table
func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&BlogPost{})
}
