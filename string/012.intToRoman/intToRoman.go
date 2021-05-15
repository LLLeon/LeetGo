package main

var valueSymbols = []struct {
	value  int
	symble string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func IntToRoman(num int) string {
	var res string

	for _, vs := range valueSymbols {
		for num >= vs.value {
			num -= vs.value
			res += vs.symble
		}
	}

	return res
}
