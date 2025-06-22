package main

import (
	"fmt"
	"slices"
)

func f(a []int) {
	b := a
	b = append(b, 10)
	fmt.Printf("%p %p %v %v\n", &a, &b, a, b)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr[2:4])
	fmt.Println(arr[0:0])

	fmt.Println(slices.Contains(arr, 3))

	data := make([][]int, 10, 50)
	fmt.Println(data, len(data))
	data[3] = nil
	fmt.Println(data, len(data))

	a1 := []int{1, 2, 3, 4, 5, 23, 2343, 123}
	fmt.Println(a1)
	a1 = append(a1, a1...)
	fmt.Println(a1)
	a2 := []int{0, 0}
	a3 := a1
	fmt.Printf("%p, %p\n", &a1, &a3)
	a3 = append(a3, 6)
	fmt.Printf("%p, %p\n", &a1, &a3)
	fmt.Println(a1, a2, a3)
	copy(a2, a1)
	fmt.Println(a1, a2)
	f(a1)
	fmt.Println(a1, a2)
}
