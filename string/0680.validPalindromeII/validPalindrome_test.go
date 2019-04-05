package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	isP := "aca"
	notP := "acba"

	if ok := IsPalindrome(isP, 0, len(isP)-1); !ok {
		t.Errorf("IsPalindrome error: want: false, get: %v\n", ok)
	} else {
		t.Log("IsPalindrome success")
	}

	if ok := IsPalindrome(notP, 0, len(notP)-1); ok {
		t.Errorf("IsPalindrome error: want: false, get: %v\n", ok)
	} else {
		t.Log("IsPalindrome success")
	}
}

func TestValidPalindrome(t *testing.T) {
	validP := "abdbdaa"

	if ok := ValidPalindrome(validP); ok {
		t.Errorf("ValidPalindrome error: want: true, get: %v\n", ok)
	} else {
		t.Log("ValidPalindrome success")
	}
}
