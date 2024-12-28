package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func Handler(ctx context.Context) {
	fmt.Println(ctx.Value("name"))
	log.Println(ctx.Deadline())
}

// 子超时时间会受到父超时时间的限制
func Controller(ctx context.Context) {
	fmt.Println(ctx.Value("name"))
	log.Println(ctx.Deadline())
	ctx = context.WithValue(ctx, "name", "ls")
	ctx, _ = context.WithTimeout(ctx, time.Second*10)
	Handler(ctx)
}

func main() {
	ctx := context.WithValue(context.Background(), "name", "zs")
	// ctx, _ = context.WithTimeout(ctx, time.Second*5) // 设置父超时时间，子超时时间受父超时时间限制
	ctx, _ = context.WithTimeout(ctx, time.Second*20)
	Controller(ctx)
}
