package main  
  
import (  
	"fmt"
)  
  
func f(done chan bool) {  
	defer func() {  
		if err := recover(); err != nil {  
			fmt.Println("Recovered in f", err)  
			done <- true // 发送信号表示已捕获panic  
		}  
	}()  
  
	a := 10
	var divisor int = 1
	divisor = 0
	_ = a / divisor
	done <- false // 如果不panic，则发送另一个信号  
}  
  
func main() {  
	done := make(chan bool)  
	go f(done)  
	if <-done {  
		fmt.Println("f function ended with panic recovered")  
	} else {  
		fmt.Println("f function ended normally (this line won't be printed)")  
	}  
	fmt.Println("main function ends")  
}