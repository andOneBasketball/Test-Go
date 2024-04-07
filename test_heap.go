package main

import (
    "fmt"
    "runtime"
	//"time"
	"reflect"
)

func main() {
    // 获取当前进程的内存使用情况
    var memStats runtime.MemStats
    runtime.ReadMemStats(&memStats)

    // 输出当前进程的内存使用情况
    fmt.Printf("Alloc = %v MiB\n", memStats.Alloc/1024/1024)
    fmt.Printf("TotalAlloc = %v MiB\n", memStats.TotalAlloc/1024/1024)
	if memStats.Sys/1024/1024 > 2 {
		fmt.Printf("%v Sys = %d MiB\n", reflect.TypeOf(memStats.Sys), memStats.Sys/1024/1024)
	}
    fmt.Printf("NumGC = %v\n", memStats.NumGC)
	//for {
	//	time.Sleep(1000 * time.Millisecond)
	//}

    // ...
}