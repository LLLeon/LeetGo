package frequent

import "testing"

func TestTopKFrequent(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	want := []int{1, 2}

	topK := TopKFrequent(nums, k)

	if !isEqual(want, topK) {
		t.Errorf("error, want: %v, get: %v\n", want, topK)
	} else {
		t.Log("success")
	}
}

func isEqual(a, b []int) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
