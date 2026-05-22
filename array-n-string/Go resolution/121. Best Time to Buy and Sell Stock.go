package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
    left := 0
    right := 1
    currentHigh := 0
    
    // 計算每段波峰與波谷之間的最大價差
    for i := 0; i < len(prices) - 1; i++ {
        temp := prices[right] - prices[left]
        if temp > 0 {
            currentHigh = max(currentHigh, temp)
        } else {
            left = right
        }
        right++
    }

    return currentHigh
}

func main() {
	nums1 := []int{7,1,5,3,6,4}
	result := maxProfit(nums1)

	fmt.Println(result)
}
