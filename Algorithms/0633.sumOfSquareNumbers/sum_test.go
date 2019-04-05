package sum

import "testing"

func TestJudgeSquareSum(t *testing.T) {
	c := 5
	want := true

	exists := JudgeSquareSum(c)
	if exists != want {
		t.Errorf("error, want: %v, get: %v\n", want, exists)
	} else {
		t.Log("success")
	}
}
