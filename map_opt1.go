package main

import "fmt"

func main() {
	var map1 map[string]*int
	d, _ := map1["abcd"]
	if d == nil {
		fmt.Println("Key not found")
	}

	a := make(map[int]int, 10)
	a[2] = 2
	a[3] = 3
	fmt.Println(len(a), a)
	delete(a, 2)
	fmt.Println(len(a), a)

	abc := map[int]int{
		10: 5,
		20: 32,
	}
	bd := map[int]int{}
	copy(bd, abc)
	bd[100] = 200
	fmt.Println(abc, bd)
}
