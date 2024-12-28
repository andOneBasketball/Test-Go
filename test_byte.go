package main

import (
	"fmt"
)

func main() {
	str := "{\"hello\":{\"a\":\"Hello, 世界!\"}}"

	// 将字符串转换为字节数组
	bytes := []byte(str)
	fmt.Printf("byte: %v, str:%s\n", bytes, str)

	data := []byte("Hello\n World")
	fmt.Println(string(data))
	fmt.Printf("%s\n", data)
}
