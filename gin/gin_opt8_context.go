package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个假的 HTTP 请求
	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBufferString(""))
	w := httptest.NewRecorder()

	// 创建一个 Gin 的上下文
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// 设置上下文中的值
	c.Set("key", "value")

	// 获取上下文中的值
	val, exists := c.Get("key")
	if exists {
		println("Key found:", val.(string)) // 输出: Key found: value
	} else {
		println("Key not found")
	}
}
