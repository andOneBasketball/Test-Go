package main

import (
	"fmt"
	"runtime"
)

func process(data []byte) {
	ref := data[:1]
	runtime.KeepAlive(ref)
}

func printMemUsage(i int) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("After %d iterations:\n", i)
	fmt.Printf("  Alloc = %v KB\n", m.Alloc/1024)
	fmt.Printf("  TotalAlloc = %v KB\n", m.TotalAlloc/1024)
	fmt.Printf("  Sys = %v KB\n", m.Sys/1024)
	fmt.Printf("  NumGC = %v\n\n", m.NumGC)
}

func main() {
	for i := 1; i <= 1000000; i++ {
		data := make([]byte, 1024)
		process(data)

		// 每 10000 次打印一次内存信息
		if i%10000 == 0 {
			printMemUsage(i)
			runtime.GC() // 手动触发 GC，看内存是否释放
		}
	}
}
