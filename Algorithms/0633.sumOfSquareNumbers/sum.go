package sum

import (
	"math"
)

func JudgeSquareSum(c int) bool {
	if c < 0 {
		return false
	}

	i := 0
	j := int(math.Sqrt(float64(c)))

	for i <= j {
		powSum := i*i + j*j
		if powSum == c {
			return true
		} else if powSum > c {
			j--
		} else {
			i++
		}
	}

	return false
}
