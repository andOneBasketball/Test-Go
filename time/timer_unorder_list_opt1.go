package main

import (
	"container/list"
	"log"
	"sync"
	"time"
)

// 插入 O(1),触发检查 O(N)

type TimerTask struct {
	expire time.Time
	job    func()
}

type UnorderedListTimer struct {
	tasks *list.List
	mu    sync.Mutex
}

func NewUnorderedListTimer() *UnorderedListTimer {
	return &UnorderedListTimer{
		tasks: list.New(),
	}
}

func (t *UnorderedListTimer) AddTimer(delay time.Duration, job func()) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.tasks.PushBack(&TimerTask{
		expire: time.Now().Add(delay),
		job:    job,
	})
}

func (t *UnorderedListTimer) Run() {
	for {
		t.mu.Lock()
		now := time.Now()
		var next *list.Element
		for e := t.tasks.Front(); e != nil; e = next {
			next = e.Next()
			task := e.Value.(*TimerTask)
			if task.expire.Before(now) {
				go task.job()
				t.tasks.Remove(e)
			}
		}
		t.mu.Unlock()
		time.Sleep(100 * time.Millisecond) // 检查间隔
	}
}

func main() {
	timer := NewUnorderedListTimer()
	go timer.Run()

	timer.AddTimer(3*time.Second, func() {
		log.Println("UnorderedListTimer: 2秒定时任务触发")
	})
	timer.AddTimer(1*time.Second, func() {
		log.Println("UnorderedListTimer: 1秒定时任务触发")
	})

	time.Sleep(4 * time.Second)
}
