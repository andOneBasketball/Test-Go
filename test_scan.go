package main

import (
	"fmt"
)

/*
fmt.Scan()  //扫描，必须所有参数都被填充后换行才结束扫描
fmt.Scanln()//扫描，但是遇到换行就结束扫描
fmt.Scanf() //格式化扫描，换行就结束
*/

func main() {
	/*
	var (
		name    string
		age     int
		//married bool
	)
	
	fmt.Scan(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
	

	//fmt.Scanf("%s %d %t", &name, &age, &married)
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("扫描结果 name:%s %d\n", name, age)
	*/

	var arr [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&arr[i][j])
		}
	}
	fmt.Print(arr)
}