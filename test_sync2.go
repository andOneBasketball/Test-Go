package main

import (
    "fmt"
    "sync"
)

func main() {
    m := sync.Map{}
    m.Store("key1", int64(1))
    m.Store("key2", int64(2))
    m.Store("key3", int64(3))

    var result string
	var count int64
    m.Range(func(key, value interface{}) bool {
        result += fmt.Sprintf("key: %s, value: %d\n", key, value)
		count += value.(int64)
		m.Store(key, int64(0))
        return true
    })

    fmt.Println(result)
	fmt.Println(count)

	value, _ := m.Load("key1")
	v := value.(int64)
	m.Store("key1", v++)

	m.Range(func(key, value interface{}) bool {
        fmt.Printf("%s %d\n", key, value)
        return true
    })
}