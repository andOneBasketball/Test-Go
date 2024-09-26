package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(5 * time.Second)
			fmt.Println(i)
			once.Do(onceBody)
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}