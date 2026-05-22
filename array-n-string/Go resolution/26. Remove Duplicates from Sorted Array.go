package main

import (
	"fmt"
)

// 本題測資已遞增排序過
func removeDuplicates(nums []int) int {
    count := 0
    temp := nums[0]

    for i := 1; i < len(nums); i++ {
		// 只要紀錄相同的元素出現次數，下次遇到不同時，把索引值減去紀錄的次數
		// 這樣的作法可以解決每刪除陣列中一個元素，就要把後面所有的元素向前推移，提高效率
        if temp != nums[i] {
            nums[i - count] = nums[i]
            temp = nums[i]
        } else {
            count++
        }
    }

    return len(nums) - count
}

func main() {
    nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 4}
	result := removeDuplicates(nums)

	fmt.Println(result)
}
