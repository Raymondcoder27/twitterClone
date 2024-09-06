package main

import (
	"example.com/twitterClone/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/tweets", controllers.GetTweets)
	r.POST("/createTweet", controllers.CreateTweet)
	r.DELETE("/tweet/:id", controllers.DeleteTweet)
}
