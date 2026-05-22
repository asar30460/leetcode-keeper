package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	var rooms int = 0 // 儲存所有過剩值（key值大於2）的總數
	var res int = 0

	var numMap = make(map[int]int)
	for i, v := range nums {
		if numMap[v] >= 2 {
			rooms++
			continue
		} else {
			nums[i-rooms] = v // 校正回歸

			numMap[v]++
			res++
		}
	}

	fmt.Print(nums, ", ")
	return res
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
}
