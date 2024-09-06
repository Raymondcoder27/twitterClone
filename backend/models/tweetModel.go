package models

import (
	"time"

	"gorm.io/gorm"
)

type Tweet struct {
	gorm.Model
	id        string `gorm:"primaryKey"`
	name      string
	handle    string
	image     string
	tweet     string
	file      string
	video     string
	comments  string
	retweets  string
	likes     string
	analytics string
	createdAt time.Time
	deletedAt gorm.DeletedAt
}
