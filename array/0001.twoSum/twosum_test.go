package sum

import (
	"testing"

	"github.com/LLLeon/LeetGo/tools"
)

func TestTwoSum(t *testing.T) {
	nums := []int{20, 30, 60, 70, 80}
	target := 100
	want := []int{0, 4}

	res1 := TwoSum(nums, target)
	if !tools.IsEqualSlice(res1, want) {
		t.Errorf("TwoSum error, want: %v, get: %v\n", want, res1)
	} else {
		t.Log("TwoSum success")
	}

	res2 := TwoSumTwoPass(nums, target)
	if !tools.IsEqualSlice(res2, want) {
		t.Errorf("TwoSumTwoPass error, want: %v, get: %v\n", want, res1)
	} else {
		t.Log("TwoSumTwoPass success")
	}

	res3 := TwoSumTwoPass(nums, target)
	if !tools.IsEqualSlice(res3, want) {
		t.Errorf("TwoSumTwoPass error, want: %v, get: %v\n", want, res1)
	} else {
		t.Log("TwoSumTwoPass success")
	}
}
