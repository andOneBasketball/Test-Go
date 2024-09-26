package main

import (
	"fmt"
	"reflect"
)

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

func main() {
	for url := range incomingURLs() {
		fmt.Println(url, reflect.TypeOf(url))
	}

	ch := incomingURLs()
	for {
		select {
		case url, ok := <- ch:
			if ok {
				fmt.Println(url, reflect.TypeOf(url))
			} else {
				fmt.Println("channel closed")
				goto OutFor
			}
		}
	}
OutFor:
	fmt.Println("main function end")
}
