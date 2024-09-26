/*
按顺序打印1~10
*/

package main

import (
	"fmt"
	"sync"
)

type ListChan struct {
	ch   chan int
	id   string
	next *ListChan
	wg   *sync.WaitGroup
}

// Print 用于输出的方法
func (l *ListChan) Print(N int, finish chan struct{}) {
	go func() {
		defer l.wg.Done()
		for {
			select {
			case num := <-l.ch:
				if num > N {
					finish <- struct{}{}
					return
				}

				fmt.Printf("%s: %d\n", l.id, num)
				l.next.ch <- num + 1

			case <-finish:
				finish <- struct{}{}
				return
			}
		}
	}()
}

// 用于构建和运行多个顺序实例的方式
func Create(M, N int) {
	ls := make([]*ListChan, M)

	wg := sync.WaitGroup{}
	finish := make(chan struct{}, M)

	for i := 0; i < M; i++ {
		ls[i] = &ListChan{
			ch: make(chan int),
			id: fmt.Sprintf("m%d", i+1),
			wg: &wg,
		}
	}

	for i := 0; i < M; i++ {
		if i == M-1 {
			ls[i].next = ls[0]
		} else {
			ls[i].next = ls[i+1]
		}
		wg.Add(1)
		ls[i].Print(N, finish)
	}
	ls[0].ch <- 1

	wg.Wait()
	for i := 0; i < M; i++ {
		close(ls[i].ch)
	}
	close(finish)
}

func main() {
	Create(5, 10)
}
