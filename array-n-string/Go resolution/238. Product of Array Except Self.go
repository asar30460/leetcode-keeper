package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	zeroIdx := []int{} // 紀錄元素值為0的位置
	maxProduct := 1 // 紀錄所有元素的乘積，除了「0」以外
	for i, v := range nums {
		if v == 0 {
			zeroIdx = append(zeroIdx, i)
			continue
		}
		maxProduct *= v
	}

	// 當陣列超過兩個或以上為0的元素，答案必定全為0
	if len(zeroIdx) >= 2 {
		return res
	}

	// 當陣列僅一個為0的元素，除該位置元素外，其餘元素都為0
	if len(zeroIdx) == 1 {
		res[zeroIdx[0]] = maxProduct
		return res
	}

	// 最大乘積除以自身O(n)
	for i, v := range nums {
		res[i] = maxProduct/v
	}

	return res
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}
