package main

import (
    "fmt"
    "sync"
)

func main() {
    m := sync.Map{}
    m.Store("key1", "value1")
    m.Store("key2", "value2")
    m.Store("key3", "value3")

	fmt.Printf("%#v\n", m)

    m.Range(func(key, value interface{}) bool {
        fmt.Printf("key: %v, value: %v\n", key, value)
        return true
    })
}