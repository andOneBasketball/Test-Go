package main

import "fmt"

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data) // `["one" "three" "three"]`

	arr1 := []int{1, 0, 2}
	arr2 := []int{1, 0, 2}
	nonemptyArr1(arr2)
	fmt.Printf("%v %v %v\n", nonemptyArr(arr1), arr1, arr2)
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonemptyArr(arr []int) []int {
	i := 0
	for _, a := range arr {
		if a != 0 {
			arr[i] = a
			i++
		}
	}
	return arr[:i]
}

func nonemptyArr1(arr []int) {
	i := 0
	for _, a := range arr {
		if a != 0 {
			arr[i] = a
			i++
		}
	}
	return
}