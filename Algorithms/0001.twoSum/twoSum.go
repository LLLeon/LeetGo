package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	for i, _ := range nums {
		for t := i + 1; t < len(nums); t++ {
			if nums[i]+nums[t] == target {
				return []int{i, t}
			}
		}
	}

	return nil
}

func twoSumTwoPass(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		m[v] = i
	}

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]

		if v, ok := m[complement]; ok && v != i {
			return []int{i, v}
		}
	}

	return nil
}

func main() {
	nums := []int{1, 2, 4, 15, 7}
	target := 19

	r := twoSum(nums, target)
	r2 := twoSumMap(nums, target)

	fmt.Println(r)
	fmt.Println(r2)
}
