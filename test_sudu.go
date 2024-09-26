package main

import (
    "fmt"
)

var bj bool

func checkPosOk(destNum int, i int, j int, arr *[9][9]int) bool {
    for x := 0; x < 9; x++ {
        if (*arr)[x][j] == destNum {
            return false
        }
    }
    for x := 0; x < 9; x++ {
        if (*arr)[i][x] == destNum {
            return false
        }
    }
    for x := i / 3 * 3; x < i / 3 * 3 + 3; x++ {
        for y := j / 3 * 3; y < j / 3 * 3 + 3; y++ {
            if (*arr)[x][y] == destNum {
                return false
            }
        }
    }
    return true
}

func dfsPos(x int, y int, arr *[9][9]int) {
    if y == 9 {
        x++
        y = 0
    }
    if x == 9 {
        bj = true
        return
    }
    if (*arr)[x][y] != 0 {
        dfsPos(x, y+1, arr)
        return
    } else {
        for i := 1; i <= 9; i++ {
            if checkPosOk(i, x, y, arr) {
                (*arr)[x][y] = i
                dfsPos(x, y+1, arr)
                if bj {
                    return
                }
                //恢复填数
                (*arr)[x][y] = 0
            }
        }
    }
    
}

func main() {
	var arr [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Scan(&arr[i][j])
		}
	}
    dfsPos(0, 0, &arr)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d", arr[i][j])
			if j < 8 {
				fmt.Print(" ")
			} else {
				fmt.Print("\n")
			}
		}
	}
}