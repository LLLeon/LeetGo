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

func twoSumOnePass(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		complement := target - nums[i]

		// 进行匹配时，第 i 个元素还没有放入 map 中，
		// 无需限制 i != index
		if index, ok := m[complement]; ok {

			// 因为刚开始遍历时还没有元素放入 map，是后面的 nums[i]
			// 跟前面放入 map 中的值匹配，所以一定是 i > index
			return []int{index, i}
		}

		m[v] = i
	}

	return nil
}

func main() {
	nums := []int{1, 4, 2, 15}
	target := 6

	r1 := twoSum(nums, target)
	r2 := twoSumTwoPass(nums, target)
	r3 := twoSumOnePass(nums, target)

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
}
