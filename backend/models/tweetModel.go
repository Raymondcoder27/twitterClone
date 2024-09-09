// models/tweet.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type Tweet struct {
	gorm.Model
	Name      string    `json:"name"`
	Handle    string    `json:"handle"`
	Image     string    `json:"image"` // New field for image
	Tweet     string    `json:"tweet"`
	File      string    `json:"file"`     // New field for file
	IsVideo   bool      `json:"is_video"` // New field for is_video
	Comments  string    `json:"comments"`
	Retweets  string    `json:"retweets"`
	Likes     string    `json:"likes"`
	Analytics string    `json:"analytics"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt
}
