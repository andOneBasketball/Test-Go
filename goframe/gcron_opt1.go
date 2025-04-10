package main

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	gcron.AddSingleton(gctx.New(), "*/1 * * * *", func(ctx context.Context) {
		g.Log().Info(gctx.New(), "每分钟执行一次")
	})

	select {} // 保持运行
}
