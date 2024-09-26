package main

import "fmt"

func main() {
    myMap := map[string]*int{}
    val := 42
    myMap["foo"] = &val

    addr := myMap["foo"]
    fmt.Println(*addr, addr) // 输出：42
}
