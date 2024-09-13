package seeders

import (
	"log"
	"time"

	"example.com/twitterClone/initializers"
	"example.com/twitterClone/models" // Adjust import path as necessary
	"gorm.io/gorm"
)

func SeedTweets() {
	db := initializers.DB
	tweets := []models.Tweet2{
		{
			Name:      "Johnni Ward",
			Handle:    "@johnniward",
			Image:     "https://randomuser.me/api/portraits/men/40.jpg",
			Tweet:     "We went rock climbing this weekend? Here is the video. Climbing is way more fun than exercising on any gym equipment...",
			File:      "/videos/Sportsman.mp4",
			IsVideo:   true,
			Comments:  "35",
			Retweets:  "54",
			Likes:     "88",
			Analytics: "81",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Josephine Perry",
			Handle:    "@josephineperry",
			Image:     "https://randomuser.me/api/portraits/women/59.jpg",
			Tweet:     "Just made a new SEO marketing video. Take a look!!! The good news is that there are several search engine optimization (SEO) tools out there...",
			File:      "/videos/Seo.mp4",
			IsVideo:   true,
			Comments:  "56",
			Retweets:  "54",
			Likes:     "78",
			Analytics: "21",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Avis Glover",
			Handle:    "@avisglover",
			Image:     "https://randomuser.me/api/portraits/men/77.jpg",
			Tweet:     "I've never seen ANYONE play guitar as good as this!!! Many guitarists neglect the theory side of things because it's perceived as boring...",
			File:      "/pics/PlayGuitar.png",
			IsVideo:   false,
			Comments:  "94",
			Retweets:  "29",
			Likes:     "33",
			Analytics: "89",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, tweet := range tweets {
		var existingTweet models.Tweet2
		// Check if tweet already exists by handle
		result := db.Where("handle = ?", tweet.Handle).First(&existingTweet)
		if result.Error == gorm.ErrRecordNotFound {
			// Tweet does not exist, create new tweet
			if err := db.Create(&tweet).Error; err != nil {
				log.Printf("Error creating tweet: %v", err)
			} else {
				log.Printf("Tweet by %s created successfully!", tweet.Handle)
			}
		} else {
			log.Printf("Tweet with handle %s already exists, skipping...", tweet.Handle)
		}
	}

	log.Println("Tweets seeding completed successfully")
}
