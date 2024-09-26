package main

import "fmt"

func main() {
    a := noReturn()
    fmt.Println(a)
}

func noReturn() (result int) {
    defer func() {
        p := recover()
        result = p.(int)
    }()
    panic(42)
}