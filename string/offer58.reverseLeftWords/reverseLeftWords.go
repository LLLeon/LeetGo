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
