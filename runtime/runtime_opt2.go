package main

import (
	"fmt"
	"runtime"
)

func getFunctionName() string {
	pc, _, _, _ := runtime.Caller(1) // 1 表示获取调用者的函数名
	return runtime.FuncForPC(pc).Name()
}

func getFunctionName2() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

func exampleFunction() {
	fmt.Println("当前函数名:", getFunctionName())
}

func main() {
	exampleFunction()
}
