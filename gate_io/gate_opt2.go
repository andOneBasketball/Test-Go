// 5 个 goroutine 并发，生成随机数量的随机数到另一个 goroutine 求和，空间复杂度 O(1)

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 初始化随机种子

	const goroutineCount = 5 // 生成随机数的 goroutine 数量

	dataChannel := make(chan int) // 用于传递随机数的通道
	var wg sync.WaitGroup

	// 启动 5 个生成随机数的 goroutine
	for i := 0; i < goroutineCount; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()            // 确保完成后调用 Done
			count := rand.Intn(10) + 1 // 随机生成数量（1 ~ 10）
			fmt.Printf("%d, count %d\n", id, count)
			for j := 0; j < count; j++ {
				num := rand.Intn(100) // 随机生成数值（0 ~ 99）
				fmt.Printf("%d Goroutine %d generated: %d\n", id, id, num)
				dataChannel <- num // 发送到通道
			}
		}(i)
	}

	// 启动一个 goroutine 用于等待所有生产者完成并关闭通道
	go func() {
		wg.Wait()          // 等待所有生成随机数的 goroutine 完成
		close(dataChannel) // 关闭通道
	}()

	// 启动求和 goroutine
	totalSum := 0
	for num := range dataChannel { // 从通道中读取数据
		totalSum += num
	}
	fmt.Printf("Total Sum: %d\n", totalSum) // 输出最终的求和结果
}
