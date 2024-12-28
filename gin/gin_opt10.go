package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MiddlewareA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("MiddlewareA: Before")
		c.Next()
		fmt.Println("MiddlewareA: After")
	}
}

func MiddlewareB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("MiddlewareB: Before")
		c.Next()
		fmt.Println("MiddlewareB: After")
	}
}

func main() {
	r := gin.Default()

	// 注册中间件
	r.Use(MiddlewareA())
	r.GET("/test1", func(c *gin.Context) {
		fmt.Println("Handler")
		c.JSON(200, gin.H{"message": "OK"})
	})

	r.Use(MiddlewareB())

	// 路由
	r.GET("/test2", func(c *gin.Context) {
		fmt.Println("Handler")
		c.JSON(200, gin.H{"message": "OK"})
	})

	r.Run(":8080")
}
