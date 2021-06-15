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

		for j < len2 && nums2[j] != curr1 {
			j++
		}

		j++
		for j < len2 && nums2[j] < curr1 {
			j++
		}

		if j == len2 {
			res[i] = -1
			continue
		}

		res[i] = nums2[j]
	}

	return res
}
