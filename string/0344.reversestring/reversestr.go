package reversestring

func ReverseString(s []byte) {
	n := len(s)
	if n <= 1 {
		return
	}

	head := 0
	tail := n - 1

	for head < tail {
		s[head], s[tail] = s[tail], s[head]
		head++
		tail--
	}
}
