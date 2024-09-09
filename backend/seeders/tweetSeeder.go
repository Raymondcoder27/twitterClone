// seeders/tweet_seeder.go
package seeders

import (
	"log"
	"time"

	"example.com/twitterClone/models" // Replace with your actual import path
	"gorm.io/gorm"
)

func SeedTweets(db *gorm.DB) {
	tweets := []models.Tweet{
		{
			Name:      "Johnni Ward",
			Handle:    "@johnniward",
			Tweet:     "We went rock climbing this weekend? Here is the video. Climbing is way more fun than exercising on any gym equipment...",
			Comments:  "35",
			Retweets:  "54",
			Likes:     "88",
			Analytics: "81",
			CreatedAt: time.Now(),
		},
		// Add more tweet entries here
	}

	for _, tweet := range tweets {
		if err := db.Create(&tweet).Error; err != nil {
			log.Fatalf("Failed to insert tweet: %v", err)
		}
	}

	log.Println("Tweets seeded successfully")
}
