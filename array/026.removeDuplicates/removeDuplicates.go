package main

func RemoveDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	counter := 1
	for p := 1; p < n; p++ {
		if nums[p] != nums[p-1] {
			nums[counter] = nums[p]
			counter++
		}
	}

	return counter
}
