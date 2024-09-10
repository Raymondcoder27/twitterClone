package models

import (
	"time"

	"gorm.io/gorm"
)

type Tweet struct {
	gorm.Model
	Name      string         `json:"name" binding:"required"` // Validation tag added
	Handle    string         `json:"handle" binding:"required"`
	Image     string         `json:"image"` // Now a simple string
	Tweet     string         `json:"tweet" binding:"required"`
	File      string         `json:"file"` // Now a simple string
	IsVideo   bool           `json:"is_video"`
	Comments  string         `json:"comments"`  // Changed back to string
	Retweets  string         `json:"retweets"`  // Changed back to string
	Likes     string         `json:"likes"`     // Changed back to string
	Analytics string         `json:"analytics"` // Could be structured further if needed
	CreatedAt int            `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
