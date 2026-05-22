package main

import (
	"fmt"
)

func canCompleteCircuit_bruteForce(gas []int, cost []int) int {
	// Time O(N^2) | Space O(1) can't pass all testcases
	diff := []int{}
	for i := 0; i < len(gas); i++ {
		diff = append(diff, gas[i]-cost[i])
	}

	// fmt.Printf("diff:%v\n", diff)

	n := len(diff)
	for i := 0; i < n; i++ {
		currGas := 0     // 計算當前油量，若加總油量低於0，表示該次迴圈沒有通過
		circuit := false // 判斷形成迴圈用

		fmt.Printf("Run%d:[", i)
		for j := i; j <= n+i; j++ {
			if currGas < 0 {
				break
			}
			if j%n == i {
				if circuit {
					fmt.Print("]\n")
					return i
				}
				circuit = true
			}
			currGas += diff[j%n]
			fmt.Printf("%d, ", currGas)
		}
		fmt.Print("]\n")
	}

	return -1
}

func canCompleteCircuit(gas []int, cost []int) int {
	totalGas, totalCost := 0, 0
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
	}

	/*
	 * 思路是我們從起點0到n-1開始記錄，但凡從該起點未到終點前就油量不足，
	 * 皆不考慮，維持這種方式一路測試到n-1...
	 * 但即便某個點油量足夠，也並不代表從該點開始能形成迴圈
	 * 所以這時候就判斷總油量是否小於總成本
	 * 如果是，表示繞不回來起點
	 * 如果否，你可能會擔心路程是不是有機會沒油
	 * 這點可以放心，如果有這個機會你的起點就不會是在這邊了
	 * 前面必定存在某個起點是可以維持旅途的油量
	 */
	start, leftGas := 0, 0 // 起點，累積油量
	for i := 0; i < len(gas); i++ {
		diff := gas[i] - cost[i]
		leftGas += diff
		if leftGas < 0 { // 如果油量不足，則更新起點為下次迭帶
			start = i + 1
			leftGas = 0
		}
	}

	if totalGas < totalCost {
		return -1
	}

	return start
}

func main() {
	// fmt.Println(canCompleteCircuit_bruteForce([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	// fmt.Println(canCompleteCircuit_bruteForce([]int{80, 2, 3, 4, 5}, []int{87, 4, 5, 1, 2}))
	// fmt.Println(canCompleteCircuit_bruteForce([]int{2, 3, 4}, []int{3, 4, 3}))

	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	fmt.Println(canCompleteCircuit([]int{7, 2, 19, 4, 5}, []int{14, 4, 16, 1, 2}))
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
}
