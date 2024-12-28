package main

import (
	"go.uber.org/zap"
)

func main() {
	// 创建一个简单的日志器（开发模式）
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // 确保日志都已写入

	// 打印一条调试级别日志
	logger.Debug("This is a debug message")

	// 打印一条结构化的日志
	logger.Info("User logged in",
		zap.String("username", "johndoe"),
		zap.Int("user_id", 42),
	)
}
