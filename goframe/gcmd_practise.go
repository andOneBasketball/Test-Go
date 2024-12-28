package main

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	Main = &gcmd.Command{
		Name:        "main",
		Brief:       "start http server",
		Description: "this is the command entry for starting your http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.BindHandler("/", func(r *ghttp.Request) {
				r.Response.Write("Hello world")
			})
			s.SetPort(8199)
			s.Run()
			return
		},
	}
)

func main() {
	Main.Run(gctx.New())
}
