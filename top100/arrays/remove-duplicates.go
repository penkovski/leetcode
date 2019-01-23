package main

import (
	"fmt"
)

func main() {

	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	n := removeDuplicates(arr)

	for i := 0; i < n; i++ {
		fmt.Println(arr[i])
	}
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	count := 1
	lastIndex, last := 0, nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > last {
			last = nums[i]
			nums[lastIndex+1] = last
			lastIndex = lastIndex + 1
			count++
		}
	}

	return count
}


