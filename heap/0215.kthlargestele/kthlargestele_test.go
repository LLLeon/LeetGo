package kthlargestele

import "testing"

func TestFindKthLargest(t *testing.T) {
	arr := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	want := 4

	if get, ok := FindKthLargest(arr, k); !ok {
		t.Logf("input parameter error, nums: %v, k: %d\n", arr, k)
	} else if get != want {
		t.Errorf("wrong result, get: %d, want: %d\n", get, want)
	} else {
		t.Log("success")
	}
}
