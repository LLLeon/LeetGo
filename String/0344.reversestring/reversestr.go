package reversestring

func ReverseString(s []byte) {
	head := 0
	tail := len(s) - 1

	for head < tail {
		s[head], s[tail] = s[tail], s[head]
		head++
		tail--
	}
}
