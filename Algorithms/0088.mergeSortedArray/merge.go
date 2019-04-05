package merge

func Merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	indexMerge := m + n - 1

	// m he n 有可能不相等，所以这里要用 ||
	for ; i >= 0 || j >= 0; indexMerge-- {

		// m 较小，继续遍历 nums2
		if i < 0 {
			nums1[indexMerge] = nums2[j]
			j--
		} else if j < 0 {
			// n 较小，继续遍历 nums1
			nums1[indexMerge] = nums1[i]
			i--
		} else if nums1[i] >= nums2[j] {
			nums1[indexMerge] = nums1[i]
			i--
		} else {
			nums1[indexMerge] = nums2[j]
			j--
		}
	}
}
