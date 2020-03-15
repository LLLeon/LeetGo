/*************************************************************************
+Author   : chenhuijia@deepglint.com
+Data     : 2020-03-14
+************************************************************************/

package power

import "testing"

func TestIsPowerOfTwo1(t *testing.T) {
	l1 := []int{1, 2, 4, 8}
	l2 := []int{3, 5, 6, 7, 9}

	for _, n := range l1 {
		if !IsPowerOfTwoI(n) {
			t.Logf("failed, %d is not power of 2.\n", n)
		} else {
			t.Logf("success, %d is power of 2.\n", n)
		}
	}

	t.Logf("//////////////////\n")

	for _, n := range l2 {
		if IsPowerOfTwoI(n) {
			t.Logf("failed, %d is power of 2.\n", n)
		} else {
			t.Logf("success, %d is not power of 2.\n", n)
		}
	}
}

func TestIsPowerOfTwo2(t *testing.T) {
	l1 := []int{1, 2, 4, 8}
	l2 := []int{3, 5, 6, 7, 9}

	for _, n := range l1 {
		if !IsPowerOfTwoII(n) {
			t.Logf("failed, %d is not power of 2.\n", n)
		} else {
			t.Logf("success, %d is power of 2.\n", n)
		}
	}

	t.Logf("//////////////////\n")

	for _, n := range l2 {
		if IsPowerOfTwoII(n) {
			t.Logf("failed, %d is power of 2.\n", n)
		} else {
			t.Logf("success, %d is not power of 2.\n", n)
		}
	}
}

func TestLog2I(t *testing.T) {
	l := []int{1, 2, 4, 8}

	want := 0
	for _, n := range l {
		power := Log2I(n)
		if power != want {
			t.Errorf("failed, want %d, get %d\n", want, power)
		} else {
			t.Logf("success, want %d, get %d\n", want, power)
			want++
		}
	}
}

func TestLog2II(t *testing.T) {
	l := []int{1, 2, 4, 8}

	want := 0
	for _, n := range l {
		power := Log2II(n)
		if power != want {
			t.Errorf("failed, want %d, get %d\n", want, power)
		} else {
			t.Logf("success, want %d, get %d\n", want, power)
			want++
		}
	}
}
