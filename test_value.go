package main
 
import (
	"fmt"
)

var a *int

func fun1() {
	if a != nil {
		fmt.Println(a)
	}
	
	*a := new int()
	a = 5
}

func main(){
	fun1()

	fun1()
}