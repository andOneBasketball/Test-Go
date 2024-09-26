package main

import (
    "fmt"
    "sync"
)

func main() {
    var (
        wg      sync.WaitGroup
        numbers []int

        mux sync.Mutex
    )

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(i int) {
            mux.Lock()
            numbers = append(numbers, i)
            mux.Unlock()

            wg.Done()
        }(i)
    }

    wg.Wait()

    fmt.Println("The numbers is", numbers)
}