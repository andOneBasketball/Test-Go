package main

import "fmt"

func main() {
	nums := []int{1, 5, 1}
	fmt.Println(nextPermutation(nums))
}

func nextPermutation(nums []int) []int {
    posX := len(nums) - 1
    var posY int = -1
    for i := len(nums) - 1; i > 0; i-- {
        if nums[i] <= nums[i-1] {
            continue
        } else {
            posY = i - 1
            for j := i; j < len(nums); j++ {
                if nums[posY] < nums[j] {
                    continue
                }
                posX = j - 1
                break
            }
            nums[posY], nums[posX] = nums[posX], nums[posY]
            break       
        }
    }

    for x, y := len(nums)-1, posY+1; x > y; x, y = x-1, y+1 {
        nums[x], nums[y] = nums[y], nums[x]
    }

    return nums
}
