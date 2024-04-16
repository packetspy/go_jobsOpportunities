package main

import "github.com/gin-gonic/gin"

func main() {
	client := gin.Default()
	client.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	client.Run() // listen and serve on 0.0.0.0:8080
}
