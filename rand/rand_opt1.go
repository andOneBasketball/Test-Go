package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Without rand.Seed:")
	for i := 0; i < 5; i++ {
		fmt.Println(rand.Intn(100)) // 不调用 rand.Seed，生成的序列固定
	}

	fmt.Println("\nWith rand.Seed:")
	rand.Seed(1) // 显式设置种子为 1
	for i := 0; i < 5; i++ {
		fmt.Println(rand.Intn(100)) // 使用种子为 1 的随机数序列
	}
}
