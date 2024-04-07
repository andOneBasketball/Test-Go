package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "-Hello, Wor-ld!"
	a := "1234"
	index := strings.ContainsRune(str, a[0])

	if index {
		fmt.Printf("字符串中第一次出现','的位置是：%v\n", index)
		//fmt.Printf("该位置之后的字符串是：%s\n", )
	} else {
		fmt.Println("字符串中没有找到','")
	}

	str = "1234"
	fmt.Printf("%v", str[1] - '0')
}
