package main

import (
    "fmt"
)

func main() {
    var cmd [10]string
	cmd[0] = "abcd"
	cmd[1] = "1234"
	for index, cmdinfo := range cmd {
		fmt.Printf("%d %s\n", index, cmdinfo)
		defer fmt.Printf("defer fun1c %d %s\n", index, cmdinfo)
	}
	fmt.Printf("%d %d\n", 5 / 3, 5 % 3)

	for i := 0; i < 10; i++ {
		defer fmt.Printf("defer fun2c %d\n", i)
	}
}