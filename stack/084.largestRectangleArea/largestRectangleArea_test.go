package main

import "testing"

func TestLargestRecAreaI(t *testing.T) {
	heights := []int{2, 1, 5, 6, 2, 3}

	maxArea := LargestRegAreaI(heights)

	want := 10
	if maxArea != want {
		t.Logf("error, want: %d, get: %d\n", want, maxArea)
	} else {
		t.Log("success")
	}
}
