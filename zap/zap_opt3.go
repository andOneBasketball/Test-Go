package main

import (
	"go.uber.org/zap"
)

func main() {
	// 创建一个简单的日志器（开发模式）
	zap.S().Info("Hello, zap!")
}
