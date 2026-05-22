package main

import (
	"fmt"
)

/*
 * 1. 字串分割
 * 2. 反轉
 */
func reverseWords(s string) string {
	s += " " // 最後一個字元若非空格，不好處理斷點
	res := ""
	splits := []string{}

	var subWord string
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if len(subWord) > 0 {
				splits = append(splits, subWord)
				subWord = ""
			}
			continue
		}
		subWord += string(s[i])
	}

	for i := len(splits) - 1; i >= 0; i-- {
		res += splits[i]
		if i != 0 {
			res += " "
		}
	}

	return res
}

func main() {
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("  hello world  "))
	fmt.Println(reverseWords("a good   example"))
	fmt.Println(reverseWords("asdasd df f"))
}
