package main

import "fmt"

type Sc struct {
	product string
	sum int
}

func sum(s []int, c chan Sc, product string) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- Sc{product: product, sum: sum}
	fmt.Printf("sum: %v\n", c)
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan Sc, 2)
	go sum(s[:len(s)/2], c, "s")
	go sum(s[len(s)/2:], c, "a")
	for {
		select {
		case x, ok := <-c:
			if !ok {
				// 通道已关闭，退出循环
				break
			}
			fmt.Println(x, "\n")
		default:
			break
		}
	}

	
	// fmt.Println(x[0].sum, y[0].sum, x[0].product, y[0].product, "\n")
	//fmt.Println(x, y, "\n")
	//fmt.Println(x.sum, y.sum, x.product, y.product, "\n")
}