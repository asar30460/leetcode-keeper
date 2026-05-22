package main

import (
	"fmt"
)

func jump(nums []int) int {
	/*
	 * 根據測資一{2, 3, 1, 1, 4}，我們預測走兩個路線的idx分別為
	 * 路線一: 0 -> 2 -> 3 -> 4
	 * 路線二: 0 -> 1 -> 4
	 * 路線一雖然開頭跑得快但他跑到終點卻更慢
	 * 要解決這個問題的話，可以設置每次跳躍的右邊界
	 * 當遍歷值碰到右邊界時表示非跳不可了
	 * 那麼此時我們的右邊界該更新成什麼呢?
	 * 舉例當idx等於0時，右邊界最多為2
	 * 在遍歷值到2之前，我們邊記錄下次跳躍最遠到位置到哪，我們稱為currFarthest
	 * 當idx等於2時，右邊界更新為currFarthest
	 * 也就是說在每次非德跳躍不可時，我們已經先找出下次的邊界了
	 * 是吧!當idx到2時，我們已經計算過路線一及路線二的下次最遠的跳躍距離了(i.e. 3及4)
	 * 所以我們只需要在碰觸到右邊界時，找出下一次的右邊界，就能以最有效率的跳法一路到終點
	 * 這便是屬於貪婪演算法的思維
	 */

	/*
	 * 注意!本題不考慮無法到達終點之測資
	 */

	jumps, rightBorder, currFarthest := 0, 0, 0
	routeIdx, letCurrFarthest := 0, currFarthest // 紀錄跳躍路線用，實際使用不須宣告
	fmt.Print("跳躍路線: [")

	for i := 0; i < len(nums)-1; i++ {
		currFarthest = max(currFarthest, i+nums[i])

		if letCurrFarthest != currFarthest {
			routeIdx = i
			letCurrFarthest = currFarthest
		}

		if i == rightBorder {
			fmt.Printf("%v -> ", routeIdx)
			jumps++
			rightBorder = currFarthest
		}
	}
	fmt.Printf("%v]，跳躍次數: ", len(nums)-1)

	return jumps
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{2, 3, 1, 3, 0, 1}))
	fmt.Println(jump([]int{2, 2, 1, 3, 0, 4}))
}
