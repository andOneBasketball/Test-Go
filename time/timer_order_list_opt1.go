package main

import (
	"container/list"
	"log"
	"sync"
	"time"
)

// 插入 O(N),触发检查 O(1)

type TimerTask struct {
	expire time.Time
	job    func()
}

type OrderedListTimer struct {
	tasks *list.List
	mu    sync.Mutex
}

func NewOrderedListTimer() *OrderedListTimer {
	return &OrderedListTimer{
		tasks: list.New(),
	}
}

func (t *OrderedListTimer) AddTimer(delay time.Duration, job func()) {
	t.mu.Lock()
	defer t.mu.Unlock()
	expire := time.Now().Add(delay)

	var prev *list.Element
	for e := t.tasks.Front(); e != nil; e = e.Next() {
		task := e.Value.(*TimerTask)
		if task.expire.After(expire) {
			break
		}
		prev = e
	}

	newTask := &TimerTask{expire: expire, job: job}
	if prev == nil {
		t.tasks.PushFront(newTask)
	} else {
		t.tasks.InsertAfter(newTask, prev)
	}
}

func (t *OrderedListTimer) Run() {
	for {
		t.mu.Lock()
		now := time.Now()
		for {
			if t.tasks.Len() == 0 {
				break
			}
			e := t.tasks.Front()
			task := e.Value.(*TimerTask)
			if task.expire.After(now) {
				break
			}
			go task.job()
			t.tasks.Remove(e)
		}
		t.mu.Unlock()
		time.Sleep(10 * time.Millisecond) // 更精确的检查间隔
	}
}

func main() {
	timer := NewOrderedListTimer()
	go timer.Run()

	timer.AddTimer(3*time.Second, func() {
		log.Println("OrderedListTimer: 2秒定时任务触发")
	})
	timer.AddTimer(1*time.Second, func() {
		log.Println("OrderedListTimer: 1秒定时任务触发")
	})

	time.Sleep(4 * time.Second)
}
