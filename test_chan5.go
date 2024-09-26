package main


import (
	"fmt"
	"time"
)

func main() {
	goroutineChan := make(chan struct{})
	go func() {
		goroutineChan <- struct{}{}
		fmt.Println("goroutine 1")
		goroutineChan <- struct{}{}
		fmt.Println("goroutine 2")
		close(goroutineChan)
	}()
	
	fmt.Println("main routine")
	time.Sleep(3 * time.Second)
	s := <- goroutineChan
	fmt.Println("1 received", s)

	time.Sleep(3 * time.Second)
	s = <- goroutineChan
	fmt.Println("2 received", s)
	time.Sleep(3 * time.Second)
}