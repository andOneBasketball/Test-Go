package main

import (
	"fmt"
	"strconv"
	"strings"
)

func joinLinks(link ...string) string {
	return strings.Join(link, ";")
}

func main() {
	result := joinLinks("https://example.com", "/path", "?query=value")
	fmt.Println(result) // 输出: https://example.com/path?query=value
	var aa int64 = 100
	fmt.Println(fmt.Sprintf("%d", aa))

	var bb string = "1233"
	num, err := strconv.Atoi(bb)
	if err != nil {
		fmt.Println("转换失败:", err)
	} else {
		fmt.Println("转换成功:", num)
	}
}
