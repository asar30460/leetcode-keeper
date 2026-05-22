package main

import (
	"fmt"
)

func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	k = k % n // 確保k不會超過n

	// 須建立part1獨立記憶體空間，否則直接對nums進行copy，part1還是會跟著改變
	part1 := make([]int, len(nums[:n-k]))
	copy(part1, nums[:n-k])

	part2 := nums[n-k:]
	copy(nums, part2)

	for i, v := range part1 {
		nums[len(part2)+i] = v
	}

	fmt.Println(nums)
}

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	rotate([]int{-1, -100, 3, 99}, 2)
}
