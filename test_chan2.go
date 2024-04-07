package main

import "fmt"

// 自定义结构体类型
type Person struct {
	Name string
	Age  int
}

func main() {
	// 创建一个通道，元素类型为自定义结构体数组
	personsChan := make(chan []Person)

	// 启动一个 goroutine 发送数据到通道
	go func() {
		persons := []Person{
			{Name: "Alice", Age: 25},
			{Name: "Bob", Age: 30},
			{Name: "Charlie", Age: 35},
		}
		personsChan <- persons
	}()

	// 从通道接收数据
	receivedPersons := <-personsChan

	// 打印接收到的数据
	fmt.Println("Received Persons:", receivedPersons)
}