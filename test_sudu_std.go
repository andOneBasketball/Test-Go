package main

import (
	"bufio"
	"fmt"
	"os"
)

var jg [][]int
var bj bool

func check(x, y, v int) bool {   //检查数字能不能填进去
	for i := 1; i <= 9; i++ {
		if jg[x][i] == v {
			return false
		}
	}
	for i := 1; i <= 9; i++ {
		if jg[i][y] == v {
			return false
		}
	}

	h, s := (x-1)/3*3+1, (y-1)/3*3+1
	for i := h; i < h+3; i++ {
		for j := s; j < s+3; j++ {
			if jg[i][j] == v {
				return false
			}
		}
	}
	return true
}

func dfs(x, y int) {
	if y == 10 {   //纵坐标跑完了横坐标+1，纵坐标回到头
		x++
		y = 1
	}
	if x == 10 {   //都跑完了，打个标记，后续的回溯不跑了
		bj = true
		return
	}
	if jg[x][y] == 0 {   //空的格子就尝试
		for i := 1; i <= 9; i++ {
			if check(x, y, i) {   //判断能不能填进去
				jg[x][y] = i   //填进去继续跑，不行就回溯回来再来一遍
				dfs(x, y+1)
				if bj {
					return
				}
				jg[x][y] = 0
			}
		}
	} else{   //有数字的格子就跳过
		dfs(x, y+1)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	jg = make([][]int, 10)
	for i := 1; i <= 9; i++ {
		jg[i] = make([]int, 10)
		for j := 1; j <= 9; j++ {
			fmt.Fscan(in, &jg[i][j])
		}
	}
	dfs(1, 1)
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			fmt.Printf("%d ", jg[i][j])
		}
		fmt.Println()
	}
}

