package main

import "fmt"

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}

func rotate(nums []int, k int)  {
	rotated := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		newindex := (i + k) % len(nums)
		rotated[newindex] = nums[i]
	}
	fmt.Println(nums)
	nums = rotated
	copy(nums, rotated)
	fmt.Println(nums)
}
