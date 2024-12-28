package main

import (
	"errors"
	"fmt"
	"sync/atomic"
)

// 乐观锁

type Data struct {
	Value   int32 // 数据值
	Version int32 // 版本号
}

// 乐观锁更新
func (d *Data) Update(newValue int32, expectedVersion int32) error {
	// 使用 CAS 实现版本号检查和更新
	fmt.Printf("Update version from %d to %d\n", expectedVersion, expectedVersion+1)
	if !atomic.CompareAndSwapInt32(&d.Version, expectedVersion, expectedVersion+1) {
		return errors.New("version mismatch, data has been modified by another goroutine")
	}
	// 如果版本号匹配，更新值
	atomic.StoreInt32(&d.Value, newValue)
	return nil
}

func main() {
	data := &Data{Value: 10, Version: 1}
	fmt.Println("Initial Data:", data)

	// 模拟并发更新
	for i := 0; i < 5; i++ {
		go func(id int) {
			expectedVersion := atomic.LoadInt32(&data.Version) // 获取当前版本号
			err := data.Update(atomic.LoadInt32(&data.Value)+1, expectedVersion)
			if err != nil {
				fmt.Printf("Goroutine %d failed: %v\n", id, err)
			} else {
				fmt.Printf("Goroutine %d updated data: %+v\n", id, data)
			}
		}(i)
	}

	// 等待一段时间，确保 Goroutine 执行完毕
	fmt.Scanln()
	fmt.Println("Final Data:", data)
}
