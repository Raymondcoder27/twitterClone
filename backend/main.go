package main

import (
	"example.com/twitterClone/controllers"
	"example.com/twitterClone/initializers"
	"example.com/twitterClone/seeders"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.MigrateDB()
}

func main() {
	seeders.SeedTweets()
	r := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}
	r.Use(cors.New(config))
	r.GET("/tweets", controllers.GetTweets)
	r.POST("/createTweet", controllers.CreateTweet)
	r.DELETE("/tweet/:id", controllers.DeleteTweet)
	r.Run(":8048")
}
