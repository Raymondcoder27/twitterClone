package main

import (
	"example.com/twitterClone/controllers"
	"example.com/twitterClone/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.MigrateDB()
}

func main() {
	r := gin.Default()
	r.GET("/tweets", controllers.GetTweets)
	r.POST("/createTweet", controllers.CreateTweet)
	r.DELETE("/tweet/:id", controllers.DeleteTweet)
	r.Run(":8089")
}
