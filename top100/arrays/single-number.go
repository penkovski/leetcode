package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))

	fmt.Println(bestSolution([]int{4, 1, 2, 1, 2}))
}

func singleNumber(nums []int) int {
	fmt.Println("[singleNumber] start")
	checked := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		checked[nums[i]]++
	}

	fmt.Println(checked)

	for k, v := range checked {
		fmt.Printf("k = %d, v = %d\n", k, v )
		if v == 1 {
			return k
		}
	}

	return 0
}

func bestSolution(nums []int) int {
	fmt.Println("[bestSolution] start")
	res := 0
	for _, num := range nums {
		fmt.Printf("res = %d, num = %d\n", res, num)
		res ^= num
		fmt.Printf("res = %d\n", res)
	}
	fmt.Println("final res =", res)
	return res
}

//4 1 2 1 2
//
//0 ^ 4 = 4 => 0000 ^ 0100 = 0100
//4 ^ 1 = 5 => 0100 ^ 0001 = 0101
//5 ^ 2 = 7 => 0101 ^ 0010 = 0111
//7 ^ 1 = 6 => 0111 ^ 0001 = 0110
