package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// maybe use auth for api

	api := r.Group("/api/v1")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Welcome to the Leaseweb Challenge API!",
			})
		})
	}

	r.Run(":8080") // Start the server on port 8080
}
