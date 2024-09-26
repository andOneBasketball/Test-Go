package main

import	"fmt"

func main() {
	arr := [4]int{2, 3, 4, 1}
	arr2 := [3]int{2, 3, 4}
	if arr == arr2 {
		fmt.Println("arr == arr2")
	}
	fmt.Println(arr, arr2)
	arr2 = append(arr2, 3)
	fmt.Println(arr2)
}



