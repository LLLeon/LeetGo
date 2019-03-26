package reversestring

import "testing"

func TestReverseString(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	result := []byte{'o', 'l', 'l', 'e', 'h'}

	ReverseString(s)
	if string(s) != string(result) {
		t.Errorf("wrong result, want: %v, get: %v\n", result, s)
	} else {
		t.Log("success")
	}
}
