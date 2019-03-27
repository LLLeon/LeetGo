package kthlargestele

import (
	"github.com/LLLeon/algoyard/algorithm/sorting/quick"
)

func FindKthLargest(arr []int, k int) (int, bool) {
	n := len(arr)
	if n == 0 || k > n {
		return 0, false
	}

	// The results are sorted from smallest to largest.
	quick.QSort(arr, n)

	return arr[n-k], true
}
