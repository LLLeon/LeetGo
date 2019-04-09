package frequent

import (
	"testing"

	"github.com/LLLeon/LeetGo/tools"
)

func TestTopKFrequent(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	want := []int{1, 2}

	topK := TopKFrequent(nums, k)

	if !tools.IsEqualSlice(want, topK) {
		t.Errorf("error, want: %v, get: %v\n", want, topK)
	} else {
		t.Log("success")
	}
}
