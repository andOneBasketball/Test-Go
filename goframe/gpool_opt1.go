package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gpool"
)

func main() {
	ctx := gctx.New()

	// 创建一个并发上限为 5 的协程池
	pool := gpool.New(5)

	// 提交 10 个任务
	for i := 0; i < 10; i++ {
		index := i // 捕获当前循环变量
		err := pool.Add(ctx, func(ctx context.Context) {
			g.Log().Infof(ctx, "任务 %d 开始执行", index)
			time.Sleep(2 * time.Second)
			g.Log().Infof(ctx, "任务 %d 执行完成", index)
		})
		if err != nil {
			g.Log().Errorf(ctx, "任务 %d 提交失败: %v", index, err)
		}
	}

	// 等待任务执行完成
	pool.CloseAndWait()
	fmt.Println("所有任务执行完毕。")
}
