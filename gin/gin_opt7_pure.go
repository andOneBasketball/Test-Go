package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 提供 unicode 实体
	r.GET("/json", func(c *gin.Context) {
		// {"html":"\u003cb\u003eHello, world!\u003c/b\u003e"}
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 提供字面字符
	r.GET("/purejson", func(c *gin.Context) {
		// {"html":"\u003Cb\u003EHello, world!\u003C/b\u003E"}
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
