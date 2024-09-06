package controllers

import (
	"errors"
	"net/http"

	"example.com/twitterClone/initializers"
	"example.com/twitterClone/models"
	"github.com/gin-gonic/gin"
)

func GetTweets(c *gin.Context) {
	var tweets []models.Tweet
	c.JSON(http.StatusOK, tweets)
}

func CreateTweet(c *gin.Context) {
	var tweet models.Tweet

	if err := initializers.DB.Create(&tweet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error creating tweet."})
	}
}

func DeleteTweet(c *gin.Context) error {
	id := c.Param("id")
	var tweet models.Tweet

	if err := initializers.DB.Where("id=?", id).First(&tweet).Error; err != nil {
		return errors.New("tweet not found")
	}

	if err := initializers.DB.Delete(&tweet).Error; err != nil {
		return err
	}

	return nil
}
