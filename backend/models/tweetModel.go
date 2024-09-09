package models

import (
	"time"

	"gorm.io/gorm"
)

type Tweet struct {
	gorm.Model
	Name      string `json:"name"`
	Handle    string `json:"handle"`
	Tweet     string `json:"tweet"`
	Comments  string `json:"comments"`
	Retweets  string `json:"retweets"`
	Likes     string `json:"likes"`
	Analytics string `json:"analytics"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
