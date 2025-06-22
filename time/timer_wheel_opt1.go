package main

import (
	"log"
	"sync"
	"time"
)

type TimeWheelTimer struct {
	slots    []map[uint64]*TimerTask
	current  int
	tick     time.Duration
	slotNum  int
	timerID  uint64
	mu       sync.Mutex
	stopChan chan struct{}
}

type TimerTask struct {
	id     uint64
	expire int
	job    func()
}

func NewTimeWheel(tick time.Duration, slotNum int) *TimeWheelTimer {
	tw := &TimeWheelTimer{
		slots:    make([]map[uint64]*TimerTask, slotNum),
		tick:     tick,
		slotNum:  slotNum,
		stopChan: make(chan struct{}),
	}
	for i := 0; i < slotNum; i++ {
		tw.slots[i] = make(map[uint64]*TimerTask)
	}
	return tw
}

func (tw *TimeWheelTimer) Start() {
	ticker := time.NewTicker(tw.tick)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			tw.tickHandler()
		case <-tw.stopChan:
			return
		}
	}
}

func (tw *TimeWheelTimer) Stop() {
	close(tw.stopChan)
}

func (tw *TimeWheelTimer) tickHandler() {
	tw.mu.Lock()
	defer tw.mu.Unlock()

	// 处理当前槽位的任务
	for id, task := range tw.slots[tw.current] {
		go task.job()
		delete(tw.slots[tw.current], id)
	}

	// 移动指针
	tw.current = (tw.current + 1) % tw.slotNum
}

func (tw *TimeWheelTimer) AddTimer(delay time.Duration, job func()) uint64 {
	tw.mu.Lock()
	defer tw.mu.Unlock()

	tw.timerID++
	ticks := int(delay / tw.tick)
	pos := (tw.current + ticks) % tw.slotNum
	task := &TimerTask{
		id:     tw.timerID,
		expire: ticks,
		job:    job,
	}
	tw.slots[pos][tw.timerID] = task
	return tw.timerID
}

func main() {
	tw := NewTimeWheel(100*time.Millisecond, 60) // 100ms精度，60个槽位(6秒一轮)
	go tw.Start()
	defer tw.Stop()

	tw.AddTimer(2*time.Second, func() {
		log.Println("TimeWheel: 2秒定时任务触发")
	})
	tw.AddTimer(1*time.Second, func() {
		log.Println("TimeWheel: 1秒定时任务触发")
	})

	time.Sleep(3 * time.Second)
}
