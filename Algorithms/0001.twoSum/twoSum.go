package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	result := make([]int, 2)

	for i, _ := range nums {
		for t := i + 1; t < len(nums); t++ {
			if nums[i]+nums[t] == target {
				result[0] = i
				result[1] = t

				break
			}
		}
	}

	return result
}

func main() {
	nums := []int{1, 2, 11, 15, 7}
	target := 9

	r := twoSum(nums, target)

	fmt.Println(r)
}
