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
	var body struct {
		Name      string `json:"name"`
		Handle    string `json:"handle"`
		Tweet     string `json:"tweet"`
		Comments  string `json:"comments"`
		Retweets  string `json:"retweets"`
		Likes     string `json:"likes"`
		Analytics string `json:"analytics"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	// var tweet models.Tweet
	tweet := models.Tweet{Name: body.Name, Handle: body.Handle, Tweet: body.Tweet, Comments: body.Comments, Retweets: body.Retweets, Likes: body.Likes, Analytics: body.Analytics}

	if err := initializers.DB.Create(&tweet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error creating tweet."})
	}

	c.JSON(http.StatusOK, "Tweet Created.")
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
