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
	name := c.PostForm("name")
	handle := c.PostForm("handle")
	tweet := c.PostForm("tweet")
	comments := c.PostForm("comments")
	retweets := c.PostForm("retweets")
	likes := c.PostForm("likes")
	analytics := c.PostForm("analytics")

	// Handle file upload
	file, _ := c.FormFile("file")
	if file != nil {
		// Save the file or process it as needed
	}

	newTweet := models.Tweet{
		Name:      name,
		Handle:    handle,
		Tweet:     tweet,
		Comments:  comments,
		Retweets:  retweets,
		Likes:     likes,
		Analytics: analytics,
		// Add file details if needed
	}

	if err := c.Bind(&newTweet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	// var tweet models.Tweet
	// tweet := models.Tweet{Name: body.Name, Handle: body.Handle, Tweet: body.Tweet, Comments: body.Comments, Retweets: body.Retweets, Likes: body.Likes, Analytics: body.Analytics}

	if err := initializers.DB.Create(&tweet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error creating tweet."})
	}

	// c.JSON(http.StatusOK, "Tweet Created.")
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
