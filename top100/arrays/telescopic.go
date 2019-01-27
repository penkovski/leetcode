package main

import "fmt"

func main() {
	fmt.Println(telescopic([]int{1, 2, 3, 2, 1, 1}))
}


func telescopic(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	if len(nums) == 1 && nums[0] != 1 {
		return false
	}

	for i := 0; i < len(nums) / 2; i++ {
		if nums[i] != i+1 {
			return false
		}

		if nums[i] != nums[len(nums)-1-i] {
			return false
		}
	}

	return true
}
