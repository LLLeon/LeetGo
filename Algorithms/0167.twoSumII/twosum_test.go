package twosum

import "testing"

func TestTwoSum(t *testing.T) {
	arr := []int{2, 7, 11, 15}
	want := []int{1, 2}

	res := TwoSum(arr, 9)
	if res[0] != want[0] || res[1] != want[1] {
		t.Errorf("get error, want: %v, get: %v\n", want, res)
	} else {
		t.Log("success")
	}
}
