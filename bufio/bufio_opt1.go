package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入一行文本：")
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	fmt.Println("你输入的是：", line)
}
