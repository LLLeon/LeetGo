package palindrome

func ValidPalindrome(s string) bool {
	n := len(s)
	if n <= 1 || n > 50000 {
		return false
	}

	i, j := 0, n-1
	for i < j {
		if s[i] != s[j] {
			return IsPalindrome(s, i+1, j) || IsPalindrome(s, i, j-1)
		}

		i++
		j--
	}

	return true
}

func IsPalindrome(s string, head, tail int) bool {
	for head < tail {
		if s[head] != s[tail] {
			return false
		}

		head++
		tail--
	}

	return true
}
