package main

import "sort"

// ugly, 超时
func ThreeSum(nums []int) [][]int {
	n := len(nums)
	empty := make([][]int, 0)
	if n == 0 {
		return empty
	}

	if n == 1 && nums[0] == 0 {
		return empty
	}

	res := make([][]int, 0)

	for i := 0; i < n; i++ {
		diff1 := -nums[i]

		for j := i + 1; j < n; j++ {
			diff2 := diff1 - nums[j]

		inner:
			for k := j + 1; k < n; k++ {
				if nums[k] == diff2 {
					tuple := []int{nums[i], nums[j], nums[k]}

					for _, tup := range res {
						if isSame(tup, tuple) {
							goto inner
						}
					}

					res = append(res, tuple)
				}
			}
		}
	}

	return res
}

func isSame(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	m := make(map[int]struct{})

	for _, v := range s1 {
		m[v] = struct{}{}
	}

	for _, v := range s2 {
		if _, ok := m[v]; !ok {
			return false
		}
	}

	return true
}

func ThreeSumI(nums []int) [][]int {
	res := [][]int{}
	counter := make(map[int]int)

	// 对数组去重并进行排序, 使接下来遍历时不会重复计算
	for _, num := range nums {
		counter[num]++
	}

	uniqNums := []int{}
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	sort.Ints(uniqNums)

	// 开始遍历去重后的数组
	for i := 0; i < len(uniqNums); i++ {
		// 统计有 3 个或以上 0 的情况
		if uniqNums[i]*3 == 0 && counter[uniqNums[i]] >= 3 {
			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i]})
		}

		for j := i + 1; j < len(uniqNums); j++ {
			// 统计两种特殊情况, 即两个相同的数字加上另一个数字的和是 0,
			if uniqNums[i]*2+uniqNums[j] == 0 && counter[uniqNums[i]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if uniqNums[i]+uniqNums[j]*2 == 0 && counter[uniqNums[j]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
			}

			c := 0 - uniqNums[i] - uniqNums[j]
			// 因为数组中的数字是递增的, 要避免 c 是 i 或 j 中的一个
			if c > uniqNums[j] && counter[c] > 0 {
				res = append(res, []int{uniqNums[i], uniqNums[j], c})
			}
		}
	}

	return res
}
