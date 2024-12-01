package main

// 读优先级读写锁

import (
	"log"
	"sync"
	"time"
)

type ReadPriorityRWLock struct {
	mu        sync.Mutex
	readCount int       // 当前活跃的读锁数量
	writeLock sync.Cond // 写锁条件变量
}

// 初始化锁
func NewReadPriorityRWLock() *ReadPriorityRWLock {
	rw := &ReadPriorityRWLock{}
	rw.writeLock.L = &rw.mu
	return rw
}

// 获取读锁
func (rw *ReadPriorityRWLock) RLock() {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	rw.readCount++
}

// 释放读锁
func (rw *ReadPriorityRWLock) RUnlock() {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	rw.readCount--
	if rw.readCount == 0 {
		rw.writeLock.Signal()
	}
}

// 获取写锁
func (rw *ReadPriorityRWLock) Lock() {
	rw.mu.Lock()
	for rw.readCount > 0 { // 等待所有读锁释放
		rw.writeLock.Wait()
	}
}

// 释放写锁
func (rw *ReadPriorityRWLock) Unlock() {
	rw.mu.Unlock()
}

func main() {
	wg := sync.WaitGroup{}
	rw := NewReadPriorityRWLock()
	count := 10
	num := 0
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			rw.RLock()
			defer rw.RUnlock()
			defer wg.Done()
			time.Sleep(time.Second)
			log.Printf("read %d %d, count %d\n", i, num, rw.readCount)
			// 读操作
			// 业务逻辑
		}()
	}
	rw.Lock()
	num++
	rw.Unlock()
	log.Printf("write %d, count %d\n", num, rw.readCount)
	wg.Wait()
}
