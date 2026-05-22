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
        arr[pivot], arr[j] = arr[j], arr[pivot]

        quicksort(arr, left, i-1)
        quicksort(arr, i+1, right)
    }
}

func majorityElement(nums []int) int {
    quicksort(nums, 0, len(nums) - 1)
    return nums[len(nums)/2]    // 排序後的中位數一定是眾數，因為題目說眾數超過一半
}

func main() {
	nums1 := []int{3,2,3}
	result := majorityElement(nums1)

	fmt.Println(result)
}
