package main

import (
	"fmt"
)

// 本題的測試都是符合羅馬數字規則
// 舉例來說，前面一個位數若小於後面位數，則減前面位數，反之加前面位數
// 但前面一個位數若小於後面位數兩階以上則不用減去
func romanToInt(s string) int {
	points := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	length := len(s)

	if length == 0 {
		return 0
	}

	sum := points[string(s[length-1])]  // 先加上最後一個位數
	for i := length - 2; i >= 0; i-- {  // 依序加上前面位數
		if points[string(s[i])] < points[string(s[i+1])] {
			sum -= points[string(s[i])]
		} else {
			sum += points[string(s[i])]
		}

	}

	return sum
}

func main() {
	result := romanToInt("MCMXCIV")

	fmt.Println(result)
}
