package main

import (
	"fmt"
	"time"
)

func twoSum(nums []int, target int) []int {
	for i := range nums {
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

		// 因为只需遍历一遍 nums，不用限制 i != index
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

	nowStart1 := time.Now()
	r1 := twoSum(nums, target)
	nowEnd1 := time.Now()
	fmt.Println("r1 excution:", nowEnd1.Sub(nowStart1))

	nowStart2 := time.Now()
	r2 := twoSumTwoPass(nums, target)
	nowEnd2 := time.Now()
	fmt.Println("r2 excution:", nowEnd2.Sub(nowStart2))

	nowStart3 := time.Now()
	r3 := twoSumOnePass(nums, target)
	nowEnd3 := time.Now()
	fmt.Println("r3 excution:", nowEnd3.Sub(nowStart3))

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
}
