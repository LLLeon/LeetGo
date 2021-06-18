package main

func NextGreaterElement(nums1, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	len1 := len(nums1)
	len2 := len(nums2)
	res := make([]int, len(nums1))

	for i := 0; i < len1; i++ {
		curr1 := nums1[i]
		j := 0

		// 找到 nums2 中等于 nums[i] 的值的索引
		for j < len2 && nums2[j] != curr1 {
			j++
		}

		// 找到 nums2 中大于 nums[i] 的值 (由于题目说明 nums1 是 nums2 的子集, 所以肯定有相同的值)
		j++
		for j < len2 && nums2[j] < curr1 {
			j++
		}

		// nums2 遍历完后没有找到满足条件的值
		if j == len2 {
			res[i] = -1
			continue
		}

		res[i] = nums2[j]
	}

	return res
}

func NextGreaterElementII(nums1, nums2 []int) []int {
	len1 := len(nums1)
	len2 := len(nums2)
	res := []int{}
	stack := []int{}
	m := make(map[int]int)

	for i := 0; i < len2; i++ {
		for len(stack) != 0 && nums2[i] > stack[len(stack)-1] {
			m[stack[len(stack)-1]] = nums2[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}

	for i := 0; i < len1; i++ {
		if v, ok := m[nums1[i]]; !ok {
			res = append(res, -1)
		} else {
			res = append(res, v)
		}
	}

	return res
}
