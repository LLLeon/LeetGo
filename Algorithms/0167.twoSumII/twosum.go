package twosum

func TwoSum(arr []int, target int) []int {
	n := len(arr)
	if n == 0 {
		return []int{}
	}

	head, tail := 0, n-1
	res := make([]int, 2)

	for head < tail {
		sum := arr[head] + arr[tail]
		if sum == target {
			res[0] = head + 1
			res[1] = tail + 1
			return res
		} else if sum > target {
			tail--
		} else {
			head++
		}
	}

	return []int{}
}
