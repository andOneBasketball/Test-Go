package main

import (
	"log"
	"sync"
	"time"
)

type HierarchicalTimerTask struct {
	expire time.Time
	job    func()
}

type HierarchicalTimeWheel struct {
	wheels     [4][]map[uint64]*HierarchicalTimerTask // 4层时间轮
	current    [4]int                                 // 各层当前指针
	resolution [4]time.Duration                       // 各层分辨率
	timerID    uint64
	mu         sync.Mutex
	stopChan   chan struct{}
}

func NewHierarchicalTimeWheel() *HierarchicalTimeWheel {
	htw := &HierarchicalTimeWheel{
		resolution: [4]time.Duration{
			100 * time.Millisecond, // 0层: 100ms
			time.Second,            // 1层: 1s
			time.Minute,            // 2层: 1m
			time.Hour,              // 3层: 1h
		},
		stopChan: make(chan struct{}),
	}

	// 初始化各层时间轮
	for i := 0; i < 4; i++ {
		htw.wheels[i] = make([]map[uint64]*HierarchicalTimerTask, 60) // 每层60个槽位
		for j := 0; j < 60; j++ {
			htw.wheels[i][j] = make(map[uint64]*HierarchicalTimerTask)
		}
	}

	return htw
}

func (htw *HierarchicalTimeWheel) Start() {
	ticker := time.NewTicker(htw.resolution[0])
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			htw.tick()
		case <-htw.stopChan:
			return
		}
	}
}

func (htw *HierarchicalTimeWheel) Stop() {
	close(htw.stopChan)
}

func (htw *HierarchicalTimeWheel) tick() {
	htw.mu.Lock()
	defer htw.mu.Unlock()

	// 处理最底层(0层)的任务
	for id, task := range htw.wheels[0][htw.current[0]] {
		go task.job()
		delete(htw.wheels[0][htw.current[0]], id)
	}

	// 推进指针并处理层级间传递
	for i := 0; i < 4; i++ {
		htw.current[i]++
		if htw.current[i] < 60 {
			break
		}
		htw.current[i] = 0
		if i < 3 {
			htw.cascade(i + 1)
		}
	}
}

func (htw *HierarchicalTimeWheel) cascade(level int) {
	for id, task := range htw.wheels[level][htw.current[level]] {
		// 重新计算位置并下放
		remaining := task.expire.Sub(time.Now())
		htw.addTimerInternal(remaining, task.job)
		delete(htw.wheels[level][htw.current[level]], id)
	}
}

func (htw *HierarchicalTimeWheel) AddTimer(delay time.Duration, job func()) uint64 {
	htw.mu.Lock()
	defer htw.mu.Unlock()
	return htw.addTimerInternal(delay, job)
}

func (htw *HierarchicalTimeWheel) addTimerInternal(delay time.Duration, job func()) uint64 {
	htw.timerID++
	expire := time.Now().Add(delay)
	task := &HierarchicalTimerTask{expire: expire, job: job}

	// 确定放入哪一层
	var level int
	for level = 0; level < 3; level++ {
		if delay < htw.resolution[level+1] {
			break
		}
	}

	// 计算槽位位置
	ticks := int(delay / htw.resolution[level])
	pos := (htw.current[level] + ticks) % 60

	htw.wheels[level][pos][htw.timerID] = task
	return htw.timerID
}

func main() {
	htw := NewHierarchicalTimeWheel()
	go htw.Start()
	defer htw.Stop()

	htw.AddTimer(2*time.Second, func() {
		log.Println("HierarchicalTimeWheel: 2秒定时任务触发")
	})
	htw.AddTimer(1*time.Second, func() {
		log.Println("HierarchicalTimeWheel: 1秒定时任务触发")
	})
	htw.AddTimer(1*time.Minute, func() {
		log.Println("HierarchicalTimeWheel: 1分钟定时任务触发")
	})

	time.Sleep(2 * time.Minute)
}
