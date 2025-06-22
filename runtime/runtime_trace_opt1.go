package main

// go tool trace trace.out
// GODEBUG=gctrace=1 go run main.go
// GODEBUG=schedtrace=1000 go run
import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

func main() {
	// 创建 trace 输出文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 启动 trace
	if err := trace.Start(f); err != nil {
		panic(err)
	}
	defer trace.Stop()

	fmt.Println("开始创建垃圾对象...")

	// 创建大量垃圾对象，强制制造 GC 压力
	for i := 0; i < 10; i++ {
		makeGarbage()
		runtime.GC()                       // 手动触发 GC
		time.Sleep(500 * time.Millisecond) // 给 GC 一些时间调度
	}

	fmt.Println("完成")
}

func makeGarbage() {
	// 创建一个临时大 slice，制造垃圾
	for i := 0; i < 10000; i++ {
		_ = make([]byte, 1024*100) // 100 KB
	}
}
