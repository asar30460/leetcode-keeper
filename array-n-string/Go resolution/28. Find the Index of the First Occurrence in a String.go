package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	if needle == "" {
		return 0
	}

	lenNeedle := len(needle)
	for i := 0; i < len(haystack)-lenNeedle+1; i++ {
		subHaystack := haystack[i : i+lenNeedle]
		// println(subHaystack)
		if subHaystack == needle {
			return i
		}
	}

	return -1
}

func main() {
	haystack := "mississippi"
	needle := "issip"

	result := strStr(haystack, needle)

	fmt.Println(result)
}
