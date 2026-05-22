package main

import (
	"fmt"
)

func quicksort(arr []int, left int, right int) {
	if left < right {
		pivot := left
		i := left
		j := right

		for i < j {
			for arr[j] > arr[pivot] && i < j {
				j--
			}

			for arr[i] <= arr[pivot] && i < j {
				i++
			}

			if i < j {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		arr[i], arr[pivot] = arr[pivot], arr[i]

		quicksort(arr, left, i-1)
		quicksort(arr, i+1, right)
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) []int {

	for i := 0; i < n; i++ {
		nums1[i+m] = nums2[i]
	}

	quicksort(nums1, 0, len(nums1)-1)

	return nums1
}

func main() {
	nums1 := []int{1}
	nums2 := []int{}
	result := merge(nums1, 1, nums2, 0)

	fmt.Println(result)
}
