package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type SpinLock struct {
	flag int32
}

// Lock 获取锁
func (s *SpinLock) Lock() {
	// 原子操作，将 flag 从 0 改置为 1
	for !atomic.CompareAndSwapInt32(&s.flag, 0, 1) {
		// 自旋等待
	}
}

// Unlock 释放锁
func (s *SpinLock) Unlock() {
	atomic.StoreInt32(&s.flag, 0)
}

func main() {
	var spinLock SpinLock

	// 示例：并发安全的计数器
	counter := 0

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("%d, 1 spinLock.flag: %d\n", i, spinLock.flag)
			spinLock.Lock()
			fmt.Printf("%d, 2 spinLock.flag: %d\n", i, spinLock.flag)
			counter++
			fmt.Printf("%d, Counter: %d\n", i, counter)
			spinLock.Unlock()
			fmt.Printf("%d, 3 spinLock.flag: %d\n", i, spinLock.flag)
		}()
	}

	// 等待一段时间，确保所有 goroutine 执行完毕
	time.Sleep(time.Second)
}
