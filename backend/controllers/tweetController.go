package controllers

import (
	"fmt"
	"net/http"

	"example.com/twitterClone/initializers"
	"example.com/twitterClone/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetTweets(c *gin.Context) {
	var tweets []models.Tweet
	if err := initializers.DB.Find(&tweets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching tweets."})
		return
	}
	c.IndentedJSON(http.StatusOK, tweets)
}

func CreateTweet(c *gin.Context) {
	// Handle file upload first
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil && err != http.ErrMissingFile {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to retrieve file: " + err.Error()})
		return
	}

	// Extract other form fields (tweet details)
	name := c.PostForm("name")
	handle := c.PostForm("handle")
	tweetText := c.PostForm("tweet")
	comments := c.PostForm("comments")
	retweets := c.PostForm("retweets")
	likes := c.PostForm("likes")
	analytics := c.PostForm("analytics")

	// Handle file upload if it exists
	var filePath string
	if file != nil {
		filePath = "./uploads/" + uuid.New().String() // Create unique file name
		if err := c.SaveUploadedFile(fileHeader, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save file"})
			return
		}
	}

	// Create the tweet object, including the file path if a file was uploaded
	tweet := models.Tweet{
		Name:      name,
		Handle:    handle,
		Tweet:     tweetText,
		Comments:  comments,
		Retweets:  retweets,
		Likes:     likes,
		Analytics: analytics,
		//FilePath:  filePath, // Add file path if there's a file
	}

	// Save the tweet to the database
	if err := initializers.DB.Create(&tweet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating tweet."})
		return
	}

	// Send the response
	c.JSON(http.StatusOK, gin.H{"message": "Tweet created successfully", "tweet": tweet})
}

func DeleteTweet(c *gin.Context) {
	id := c.Param("id")
	var tweet models.Tweet

	if err := initializers.DB.Where("id=?", id).First(&tweet).Error; err != nil {
		fmt.Println("tweet not found")
		return
	}

	if err := initializers.DB.Delete(&tweet).Error; err != nil {
		fmt.Printf("tweet not found %v", err)
		return
	}

}
