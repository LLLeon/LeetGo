package tools

import "testing"

func TestIsEqualSlice(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	c := []int{3, 2, 1}
	d := []int{}
	var e []int
	f := make([]int, 0)

	if isEqual := IsEqualSlice(a, b); !isEqual {
		t.Errorf("error comparing %v with %v\n", a, b)
	} else {
		t.Log("success")
	}

	if isEqual := IsEqualSlice(a, c); isEqual {
		t.Errorf("error comparing %v with %v\n", a, c)
	} else {
		t.Log("success")
	}

	if isEqual := IsEqualSlice(c, d); isEqual {
		t.Errorf("error comparing %v with %v\n", c, d)
	} else {
		t.Log("success")
	}

	if isEqual := IsEqualSlice(a, e); isEqual {
		t.Errorf("error comparing %v with %v\n", a, e)
	} else {
		t.Log("success")
	}

	if isEqual := IsEqualSlice(d, e); isEqual {
		t.Errorf("error comparing %v with %v\n", d, e)
	} else {
		t.Log("success")
	}

	if isEqual := IsEqualSlice(d, f); !isEqual {
		t.Logf("d addr: %p, f addr: %p\n", d, f)
		t.Errorf("error comparing %v with %v\n", d, f)
	} else {
		t.Log("success")
	}

	if isEqual := IsEqualSlice(e, f); isEqual {
		t.Errorf("error comparing %v with %v\n", e, f)
	} else {
		t.Log("success")
	}
}
