package main

import (
	"fmt"
	"os"
)

func main() {
	envVar := os.Getenv("MY_VAR")
	fmt.Println("MY_VAR =", envVar)
}
