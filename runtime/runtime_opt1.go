package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {

			}
		}()
	}
	//wg.Wait()
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Number of Go routines:", runtime.NumGoroutine())
	wg.Wait()
}
