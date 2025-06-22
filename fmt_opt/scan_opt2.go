package main

import (
	"fmt"
)

func main() {
	var row, line, data, posI, posJ int
	fmt.Scan(&row, &line)
	arr := make([][]int, row)
	for i := range arr {
		arr[i] = make([]int, line)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < line; j++ {
			fmt.Scan(&data)
			arr[i][j] = data
		}
	}
	fmt.Scan(&posI, &posJ)
	fmt.Println(arr, posI, posJ, row, line)
}
