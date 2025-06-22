package main

import (
	"fmt"
)

func main() {
	var row, col int
	fmt.Scan(&row, &col)
	data := ""
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Scan(&data)
			fmt.Println(data)
		}
	}
	fmt.Scan(&data)
	fmt.Println(data)
}
