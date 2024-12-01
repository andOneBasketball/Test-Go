package main

import (
	"sync"
)

type FairRWLock struct {
	mu         sync.Mutex
	readers    int  // 当前活跃的读锁数量
	writerWait bool // 是否有写锁等待
	readWait   sync.Cond
	writeWait  sync.Cond
}

// 初始化公平锁
func NewFairRWLock() *FairRWLock {
	lock := &FairRWLock{}
	lock.readWait.L = &lock.mu
	lock.writeWait.L = &lock.mu
	return lock
}

// 获取读锁
func (rw *FairRWLock) RLock() {
	rw.mu.Lock()
	defer rw.mu.Unlock()

	for rw.writerWait { // 如果有写锁等待，阻塞读锁
		rw.readWait.Wait()
	}
	rw.readers++
}

// 释放读锁
func (rw *FairRWLock) RUnlock() {
	rw.mu.Lock()
	defer rw.mu.Unlock()

	rw.readers--
	if rw.readers == 0 {
		rw.writeWait.Signal() // 唤醒写锁
	}
}

// 获取写锁
func (rw *FairRWLock) Lock() {
	rw.mu.Lock()
	defer rw.mu.Unlock()

	rw.writerWait = true
	for rw.readers > 0 { // 等待所有读锁释放
		rw.writeWait.Wait()
	}
}

// 释放写锁
func (rw *FairRWLock) Unlock() {
	rw.mu.Lock()
	defer rw.mu.Unlock()

	rw.writerWait = false
	rw.readWait.Broadcast() // 唤醒所有读锁
	rw.writeWait.Signal()   // 唤醒下一个写锁
}
