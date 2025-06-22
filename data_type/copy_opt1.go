package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := make([]int, len(arr1), len(arr1))
	copy(arr2, arr1)
	arr2[0] = 10
	fmt.Println(arr1, arr2)
}
