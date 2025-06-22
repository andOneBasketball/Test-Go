package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("go end")
				return
			default:
				fmt.Println("go running")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}
