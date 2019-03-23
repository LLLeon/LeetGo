package palindrome

import (
	"bytes"
	"strings"
)

func IsPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	var buf bytes.Buffer

	for i := 0; i <= len(s)-1; i++ {
		if IsAlphanum(s[i]) {
			ele := strings.ToLower(string(s[i]))
			buf.WriteString(ele)
		}
	}

	cleanStr := buf.String()
	head, tail := 0, len(cleanStr)-1

	for head < tail {
		if cleanStr[head] == cleanStr[tail] {
			head++
			tail--
		} else {
			return false
		}
	}

	return true
}

func IsPalindrome2(s string) bool {
	if len(s) == 0 {
		return true
	}

	head, tail := 0, len(s)-1

	for head < tail {
		for head < tail && !IsAlphanum(s[head]) {
			head++
		}

		for head < tail && !IsAlphanum(s[tail]) {
			tail--
		}

		if strings.ToLower(string(s[head])) != strings.ToLower(string(s[tail])) {
			return false
		}

		head++
		tail--
	}

	return true
}

func IsAlphanum(ele uint8) bool {
	return ('a' <= ele && ele <= 'z') || ('A' <= ele && ele <= 'Z') || ('0' <= ele && ele <= '9')
}
