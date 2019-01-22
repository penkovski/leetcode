package main

import "fmt"

func main() {
	nums := twosum([]int{-1, -2, -3, -4, -5}, -8)
	fmt.Println(nums)
}

func twosum(nums []int, target int) []int {

	checked := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := checked[nums[i]]; ok {
			continue
		}

		for j := i+1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}

		checked[nums[i]] = i
	}

	return nil
}

//11 2 15 7

