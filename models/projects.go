package models

import (
	"time"
)

type Projects struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string
	Description string
	Link        string
	ImageURL    string
	CreatedAt   time.Time
}


// our service structure model
type Services struct{
	ID  uint `gorm:"primaryKey"`
	Title string
	Description string
	ImageURL string
	CreatedAt time.Time

}
