package models

import (
	"time"

	"gorm.io/gorm"
)

type Tweet struct {
	gorm.Model
	Id        string `gorm:"primaryKey"`
	Name      string
	Handle    string
	Image     string
	Tweet     string
	File      string
	Video     string
	Comments  string
	Retweets  string
	Likes     string
	Analytics string
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
