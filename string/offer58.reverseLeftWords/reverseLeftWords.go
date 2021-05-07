package main

func ReverseLeftWords(s string, n int) string {
	if len(s) == 0 || n <= 0 {
		return s
	}

	if len(s) < n {
		return s
	}

	left := s[:n]
	right := s[n:]

	return right + left
}

func ReverseLeftWordsII(s string, n int) string {
	if len(s) == 0 || n <= 0 {
		return s
	}

	if len(s) < n {
		return s
	}

	orgLen := len(s)
	s += s

	return s[n : orgLen+n]
}
