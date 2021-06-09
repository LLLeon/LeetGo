package main

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
