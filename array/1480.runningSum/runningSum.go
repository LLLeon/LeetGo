package main

// 时间复杂度 O(N2), 空间复杂度 O(N).
func RunningSum(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	res := []int{}

	for i := range nums {
		frontPart := nums[:i+1]
		sum := 0

		for _, v := range frontPart {
			sum += v
		}

		res = append(res, sum)
	}

	return res
}

// 时间复杂度 O(N), 空间复杂度 O(N).
func RunningSumII(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	res := []int{nums[0]}

	for i := 1; i < len(nums); i++ {
		res = append(res, res[i-1]+nums[i])
	}

	return res
}
