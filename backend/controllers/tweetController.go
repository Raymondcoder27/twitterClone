package controllers

import (
	"fmt"
	"net/http"

	"example.com/twitterClone/initializers"
	"example.com/twitterClone/models"
	"github.com/gin-gonic/gin"
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
	var requestData models.Tweet

	// Bind the JSON request payload to requestData
	if err := c.Bind(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	// Handle file upload
	file, err := c.FormFile("file")
	if err != nil {
		// If no file, just process the tweet text
		c.JSON(http.StatusOK, gin.H{
			"message": "Tweet created successfully without a file",
			"tweet":   requestData,
		})
		return
	}

	// If a file is uploaded, handle it
	path := "./uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "file saving failed"})
		return
	}

	newTweet := models.Tweet{
		Name:      requestData.Name,
		Handle:    requestData.Handle,
		Tweet:     requestData.Tweet,
		Comments:  requestData.Comments,
		Retweets:  requestData.Retweets,
		Likes:     requestData.Likes,
		Analytics: requestData.Analytics,
		// Add file details if needed
	}

	if err := initializers.DB.Create(&newTweet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating tweet."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tweet created successfully"})
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
