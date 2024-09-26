package main

import (
	"fmt"
)

func main() {
	fmt.Println("a return:", a()) // 打印结果为 return: 0
	fmt.Println("b return:", b())
	fmt.Println("c return:", *(c()))
}

func a() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()
	return i
}

func b() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()
	return i // 或者直接 return 效果相同
}

func c() *int {
	var i int
	defer func() {
		i++
		fmt.Println("c defer2:", i) // 打印结果为 c defer: 2
	}()
	defer func() {
		i++
		fmt.Println("c defer1:", i) // 打印结果为 c defer: 1
	}()
	return &i
}