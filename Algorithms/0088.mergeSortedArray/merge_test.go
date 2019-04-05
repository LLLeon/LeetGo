package merge

import "testing"

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	want := []int{1, 2, 2, 3, 5, 6}

	Merge(nums1, m, nums2, n)

	if len(nums1) != len(want) {
		t.Errorf("the number of elements does not match, want: %d, get: %d\n",
			len(want), len(nums1))
	}

	// 防止越界
	nums1 = nums1[:len(want)]
	for i, v := range want {
		if nums1[i] != v {
			t.Errorf("element does not match: want: %d, get: %d\n", v, nums1[i])
		}
	}

	t.Log("success")
}
