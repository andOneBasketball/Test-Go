package main

import (
	"fmt"
	"time"
)

type LogPrintData struct {
	count int64
	timestamp time.Time
}

type Person struct {
    Name string
    Age  int
}

func main() {
	logPrintCount := make(map[string]LogPrintData)
	logPrintCount["s"] = LogPrintData{count: 0, timestamp: time.Now()}
	data := logPrintCount["s"]
	data.count += 1
	logPrintCount["s"] = data

	fmt.Printf("%#v\n", logPrintCount)

	p1 := Person{Name: "Alice", Age: 20}
    p2 := p1 // 直接将p1赋值给p2
    fmt.Println(p2) // 输出: {Alice 20}
}
