package main

//指定数量协程

import (
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/google/uuid"
)

func main() {
	var wg sync.WaitGroup
	maxGoroutines := 4
	goroutineChan := make(chan struct{}, maxGoroutines)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		goroutineChan <- struct{}{}
		go func(i int) {
			defer func() {
				<-goroutineChan
				wg.Done()
			}()
			// 协程的具体逻辑
			uuid := uuid.NewString()
			fmt.Printf("start goroutine ID: %v %s\n", reflect.TypeOf(uuid), uuid)
			//time.Sleep(time.Duration(i) * 2000 * time.Millisecond)
			fmt.Printf("end goroutine ID(%v) %d %v\n", uuid, i, time.Now())
		}(i)
	}

	wg.Wait()
}
