package models

import (
	"time"
)

type Projects struct {
    ID          uint      `gorm:"primaryKey"`
    Title       string    `gorm:"not null"`
    Slug        string    `gorm:"uniqueIndex;not null"`
    Description string    `gorm:"type:text"`
    ImageURL    string    `gorm:"type:text"`       // project cover image
    Link        string    `gorm:"type:text"`       // external link (e.g., GitHub, live demo)
    Category    string    `gorm:"size:100"`        // e.g. SaaS, Mobile, AI, etc.
    Tags        string    `gorm:"type:text"`       // comma-separated list, or use separate table if needed
    Featured    bool      `gorm:"default:false"`   // highlight in frontend (hero/banner)
    Published   bool      `gorm:"default:true"`    // control visibility
    
    // Relationship
	TechStacks []TechStack `gorm:"many2many:project_techstacks;"`

    // Metadata
    CreatedAt   time.Time
    UpdatedAt   time.Time
}


// Service model (renamed to singular for convention)
// Service model
type Services struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"size:200;not null"`
	Slug        string    `gorm:"uniqueIndex;not null"`
	Description string    `gorm:"type:text"`
	ImageURL    string    `gorm:"type:text"`

	// Categorization & metadata
	Category  string `gorm:"size:100"`
	Tags      string `gorm:"type:text"` // e.g. "cloud,hosting,servers"
	Featured  bool   `gorm:"default:false"`
	Published bool   `gorm:"default:true"`

	// Author (admin who created it)
	AuthorID *uint
	Author   *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	// TechStack Many-to-Many relationship
	TechStacks []TechStack `gorm:"many2many:service_techstacks;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}



// my tech stack structure model
// TechStack model (shared between Projects & Services)
type TechStack struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:100;uniqueIndex"`
	IconURL   string    `gorm:"type:text"` // optional: icon/logo for frontend display
	CreatedAt time.Time
	UpdatedAt time.Time

	// Reverse relations (optional)
	Projects []Projects `gorm:"many2many:project_techstacks;"`
	Services []Services `gorm:"many2many:service_techstacks;"`
}