package main

import (
    "fmt"
)

func main() {
    var str string
    for {
        n, _ := fmt.Scan(&str)
        if n == 0 {
            break
        } else {
            m := make(map[uint8]bool)
            num := 0
            for i := 0; i < len(str); i++ {
                if _, exist := m[str[i]]; exist {
                    continue
                }
                m[str[i]] = true
                num++
            }
            fmt.Printf("%d\n", num)
        }
    }
}