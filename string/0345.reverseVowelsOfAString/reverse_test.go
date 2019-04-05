package reverse

import "testing"

func TestReverseVowels(t *testing.T) {
	s := "hello"
	want := "holle"

	getStr := ReverseVowels(s)
	if getStr != want {
		t.Errorf("error, want: %v, get: %v\n", want, getStr)
	} else {
		t.Log("success")
	}
}
