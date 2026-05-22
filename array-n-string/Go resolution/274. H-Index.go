package main

import (
	"fmt"
	"sort"
)

// h-index 找出給定作者的第i篇論文至少要有i次引用的最大值
func hIndex(citations []int) int {
	// 降冪排序
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))

	var i int = 0
	for ; i < len(citations); i++ {
		// 當遍歷到不符合條件時，由於迴圈是從零開始，直接回傳i剛好一正一負抵銷
		// 判斷是要使用小於等於而非僅小於，同樣，因為迴圈從零開始
		// 所以發現等於時，實際上引用次數已經是i+1了
		if citations[i] <= i {
			break
		}
	}

	return i
}

func main() {
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
	fmt.Println(hIndex([]int{1, 3, 1}))
}
