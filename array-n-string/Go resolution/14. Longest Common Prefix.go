package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
    
    base := strs[0]	// 以str第一個元素為基準
	lcp_len := len(base)  // 紀錄目前最長相同位數
	for _, s := range strs {
		i := 0

		// 比較與其它元素的最高相同位數
		for ; i < len(s) && i < len(base) && s[i] == base[i]; i++ {
		}

		lcp_len = min(lcp_len, i)	// 更新最長相同位數
	}

	return base[0:lcp_len]	// 擷取最長相同位數的字母
}

func main() {
	strs := []string{"flower","flow","flight"}
	result := longestCommonPrefix(strs)

	fmt.Println(result)
}
