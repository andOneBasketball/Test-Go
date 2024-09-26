package main

import (
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

func main() {
    var (
        mux    sync.Mutex
        state1 = map[string]int{
            "a": 65,
        }
        muxTotal uint64

        rw     sync.RWMutex
        state2 = map[string]int{
            "a": 65,
        }
        rwTotal uint64

        mux2    sync.Mutex
        state3 = map[string]int{
            "a": 65,
        }
        mux2Total uint64
    )

    for i := 0; i < 10; i++ {
        go func() {
            for {
                mux2.Lock()
				// 读性能
                // _ = state1["a"]
				// 写性能
				state3["a"] = 66
                mux2Total++
                mux2.Unlock()
                // atomic.AddUint64(&muxTotal, 1)
            }
        }()
    }

    for i := 0; i < 10; i++ {
        go func() {
            for {
                mux.Lock()
				// 读性能
                // _ = state1["a"]
				// 写性能
				state1["a"] = 66
                // muxTotal++
                mux.Unlock()
                atomic.AddUint64(&muxTotal, 1)
            }
        }()
    }

    for i := 0; i < 10; i++ {
        go func() {
            for {
				//rw.RLock()
                rw.Lock()
                //_ = state2["a"]
				// _ = state1["a"]
				// 写性能
				state2["a"] = 66
				// rw.RUnlock()
                rw.Unlock()
                atomic.AddUint64(&rwTotal, 1)
            }
        }()
    }

    time.Sleep(time.Second)

    fmt.Println("sync.Mutex readOps is", muxTotal, " ", mux2Total)
    fmt.Println("sync.RWMutex readOps is", rwTotal)
}