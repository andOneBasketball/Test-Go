package main

import (
	"github.com/gin-gonic/gin"
)

// curl -v -X POST -F "message=login" -F "nick=password" http://localhost:8080/form_post
// curl -v -X POST -F "message=login" http://localhost:8080/form_post

func main() {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.Run(":8080")
}
