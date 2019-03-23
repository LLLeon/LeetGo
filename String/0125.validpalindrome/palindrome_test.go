package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	s := "A man, a plan, a canal: Panama"

	if ok := IsPalindrome(s); !ok {
		t.Error("wrong result")
	} else {
		t.Log("success")
	}
}

func TestIsPalindrome2(t *testing.T) {
	s := "A man, a plan, a canal: Panama"

	if ok := IsPalindrome2(s); !ok {
		t.Error("wrong result")
	} else {
		t.Log("success")
	}
}

func TestIsAlphanum(t *testing.T) {
	s := "A man, a plan, a canal: Panama"

	for i := 0; i <= len(s)-1; i++ {
		if !IsAlphanum(s[i]) {
			t.Logf("%v is not alphanum:\n", string(s[i]))
		}
	}

	t.Log("success")
}
