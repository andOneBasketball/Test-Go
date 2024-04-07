package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("%s~%s\n", t.AddDate(0, 0, -1).Format("2006.01.02"), t.Format("01.02"))
	fmt.Println(t.Format("20060102150405"))

	ticker := time.Tick(1 * time.Second)
	var msgCountStr string
	var allProductMsgCount int64
	for {
		select {
		case <-ticker:
			fmt.Printf("%s %d\n", msgCountStr, allProductMsgCount)
			msgCountStr += "12345"
			allProductMsgCount += 2
		}
	}
}