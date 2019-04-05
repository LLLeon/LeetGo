package reverse

func ReverseVowels(s string) string {
	if s == "" || s == " " {
		return s
	}

	vowels := map[uint8]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}

	n := len(s)
	i, j := 0, n-1
	res := []rune(s)

	for i <= j {
		ci := s[i]
		cj := s[j]

		if _, ok := vowels[ci]; !ok {
			res[i] = rune(ci)
			i++
		} else if _, ok := vowels[cj]; !ok {
			res[j] = rune(cj)
			j--
		} else {
			res[i] = rune(cj)
			res[j] = rune(ci)
			i++
			j--
		}
	}

	return string(res)
}
