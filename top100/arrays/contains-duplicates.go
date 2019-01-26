package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1, 4}))
}

func containsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	checked := make(map[int]bool)
	checked[nums[0]] = true
	for i := 1; i < len(nums); i++ {
		if _, ok := checked[nums[i]]; ok {
			return true
		}
		checked[nums[i]] = true
	}
	return false
}
