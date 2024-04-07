package main

import (
	"fmt"
	"time"
)


func main() {
	a := 7
	fmt.Println(time.Now().Unix(), time.Now())
	fmt.Println(int(time.Now().AddDate(0, 0, -a).Unix()), time.Now().AddDate(0, 0, -7))

	identifyStartTime := time.Now()
	fmt.Println(identifyStartTime.Unix())
	if time.Now().Unix() - identifyStartTime.Unix() > 0 {
		fmt.Println("abce")
	}

	bulkOpr := make(map[string][]interface{})
	opr := []interface{}{
		"abcdsf",
	}
	bulkOpr["s"] = append(bulkOpr["s"], opr)
	fmt.Printf("%v", bulkOpr)

	var ace int
	ace = 3
	t := time.Now().Add(-time.Hour * time.Duration(ace)).Unix()
    fmt.Println(t)
}
