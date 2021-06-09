package main

import "sort"

func SmallerNumbersThanCurrent(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	res := make([]int, 0)

	for _, curr := range nums {
		counter := 0
		for _, v := range nums {
			if v < curr {
				counter++
			}
		}
		res = append(res, counter)
	}

	return res
}

func SmallerNumbersThanCurrentII(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	res := make([]int, 0)
	oldNums := make([]int, len(nums))
	copy(oldNums, nums)
	vIdxMap := make(map[int]int)

	// 从小到大排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] <= nums[j]
	})

	// 注意这里, 有相同的数字时, 就不用再记录到 map 中了, 否则会以最后记录的数字的索引为结果, 这样不准确.
	for i, v := range nums {
		if _, ok := vIdxMap[v]; !ok {
			vIdxMap[v] = i
		}
	}

	// 遍历原始数组, 按元素原位置统计结果.
	for _, v := range oldNums {
		res = append(res, vIdxMap[v])
	}

	return res
}
