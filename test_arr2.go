package	main

import "fmt"

func f(arr *[]int) {
	fmt.Println((*arr)[1], (*arr)[2])
}

func main() {
	foo := []int{0, 0, 0, 42, 100}
	bar := foo[1:4:4]
	fmt.Println("bar:", bar)
	bar = append(bar, 99)
	fmt.Println("foo:", foo) // foo: [0 0 0 42 100]
	fmt.Println("bar:", bar) // bar: [0 0 42 99]
	f(&foo)
}

