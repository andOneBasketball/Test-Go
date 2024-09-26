package main

/*
有 N
 件物品和一个容量是 V
 的背包。每件物品只能使用一次。

第 i
 件物品的体积是 vi
，价值是 wi
。

求解将哪些物品装入背包，可使这些物品的总体积不超过背包容量，且总价值最大。
输出最大价值。

输入格式
第一行两个整数，N，V
，用空格隔开，分别表示物品数量和背包容积。

接下来有 N
 行，每行两个整数 vi,wi
，用空格隔开，分别表示第 i
 件物品的体积和价值。

输出格式
输出一个整数，表示最大价值。

数据范围
0<N,V≤1000

0<vi,wi≤1000
输入样例
4 5
1 2
2 4
3 4
4 5
输出样例：
8

f(n, v) = max(f(n-1, v), f(n-1, v-c[n]) + w[n])

2 5
4 6
5 4

f[1][5] = max(f[0][5], f[0][0] + 4) = max(0, 0 + 6) = 6
f[0][5] = max(f[0][1] + 6)

*/


import (
	"fmt"
	"os"
)

func max(a, b int) (int) {
	if a > b {
		return a
	}
	return b
}

func backpack(n int, v int, arr [][2]int, fn *[][]int, w *[]int) {
	sum := 0
	for i := n; i > 0; i-- {
		fmt.Println("i:", i, "v:", v, "arr:", arr[i], "fn:", (*fn)[i])
		if v - arr[i][0] < 0 {
			(*fn)[i][v] = (*fn)[i-1][v]
		} else {
			(*fn)[i][v] = max((*fn)[i-1][v], (*fn)[i-1][v-arr[i][0]] + arr[i][1])
			if (*fn)[i-1][v] <= (*fn)[i-1][v-arr[i][0]] + arr[i][1] {
				v = v - arr[i][0]
				sum += arr[i][1]
			}
		}
		backpack(i, v, arr, fn, w)
	}
	*w = append(*w, sum)
}

func main() {
	file, err := os.Open("input.txt")
	if err!= nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	os.Stdin = file

	var n, v int
	fmt.Scan(&n, &v)
	arr := make([][2]int, n + 1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&arr[i][0], &arr[i][1])
	}
	fn := make([][]int, n)
	for i := 1; i <= n; i++ {
		fn[i] = make([]int, v + 1)
	}
	var w []int
	backpack(n, v, arr, &fn, &w)
	
	maxValue := 0
	for _, j := range w {
		if maxValue < j {
			maxValue = j
		}
	}

	fmt.Println(maxValue)
}


