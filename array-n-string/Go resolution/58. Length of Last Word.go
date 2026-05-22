package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
    last_sum :=0    // 紀錄最後一個非空白的字串長度
    sum := 0    // 記錄所有字串的長度（包含空白）

    // 判斷每個字元是否為空格
    for i :=0; i<len(s); i++ {
        if string(s[i]) != " " {
            sum++
            last_sum = sum
        } else {
            sum = 0
        }
    }

    if sum == 0 {
        return last_sum
    }

    return sum
}

func main() {
	strs := "   fly me   to   the moon  "
	result := lengthOfLastWord(strs)

	fmt.Println(result)
}
