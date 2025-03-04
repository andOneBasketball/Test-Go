package main

import (
	"fmt"
	"log"
	"time"
)

func f(i *int, beginTime time.Time) {
	syncRemarkCostTime := time.Now().Sub(beginTime).Seconds()
	fmt.Printf("%d %f\n", *i, syncRemarkCostTime)
	*i += 5
	*i++
}

func f2() int64 {
	a := int64(0)
	a++
	return a
}

func main() {
	i := 4
	f(&i, time.Now())
	fmt.Printf("%d\n", i)

	var aa int64 = 100
	log.Println(float64(aa))
}
