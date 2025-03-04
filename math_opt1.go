package main

import (
	"fmt"
	"math"
)

func main() {
	num := 2.5
	rounded := math.Round(num) // 四舍五入
	fmt.Println(rounded)       // 输出 3
}
