package main

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
