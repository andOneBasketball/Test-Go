package main

import "fmt"

func f1(arr *[4]int) {
	arr[0] = 100
}

func main() {
	arr := [4]int{2, 3, 4, 1}
	fmt.Println(arr)
	f1(&arr)
	fmt.Println(arr)
	/*
		arr2 := [3]int{2, 3, 4}
		if arr == arr2 {
			fmt.Println("arr == arr2")
		}
		fmt.Println(arr, arr2)
		arr2 = append(arr2, 3)
		fmt.Println(arr2)
	*/
}
