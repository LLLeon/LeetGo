package word

import "testing"

func TestFindLongestWord(t *testing.T) {
	s := "abpcplea"
	d := []string{"ale", "apple", "monkey", "plea"}
	want := "apple"

	get := FindLongestWord(s, d)
	if get != want {
		t.Errorf("FindLongestWord error, want: %v, get: %v\n", want, get)
	} else {
		t.Log("FindLongestWord success")
	}
}

func TestIsContain(t *testing.T) {
	src := "abcdef"
	sub := "bdf"

	if ok := IsContain(src, sub); !ok {
		t.Errorf("IsContain error, want: true, get: %v\n", ok)
	} else {
		t.Log("IsContain success")
	}
}
