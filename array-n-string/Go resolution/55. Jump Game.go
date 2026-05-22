package main

import (
	"fmt"
)

func canJump(nums []int) bool {
	/*
	 * 設定變數(maxLen)記錄從idx[0]開始最大累積跳躍步數
	 * 遍歷整個陣列，當遍歷到一個點時發現大於maxLen
	 * 表示所有的路徑都無法到達該點
	 */
	maxLen := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if i > maxLen {
			return false
		}

		/* 如果其中一個元素已經可以到達終點直接回傳
		 * 但是會不會更有效率取決於測試集
		 * 萬一能達到終點的都是接近終點的值
		 * 變成每次迭代都多一次這個判斷，整體來看反而更沒效率
		 */
		if maxLen >= n-1 {
			return true
		}

		maxLen = max(maxLen, i+nums[i])
	}

	return true
}

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{2, 3, 1, 3, 0, 1}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
