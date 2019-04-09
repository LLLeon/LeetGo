package sum

func TwoSum(nums []int, target int) []int {
	for i := range nums {
		for t := i + 1; t < len(nums); t++ {
			if nums[i]+nums[t] == target {
				return []int{i, t}
			}
		}
	}

	return nil
}

func TwoSumTwoPass(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		m[v] = i
	}

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]

		if index, ok := m[complement]; ok && index != i {
			return []int{i, index}
		}
	}

	return nil
}

func TwoSumOnePass(nums []int, target int) []int {
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
