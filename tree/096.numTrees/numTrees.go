package main

// 这个解法在 leetcode 提交时超时, 因为有重复子问题, 时间复杂度爆炸.
func NumTrees(n int) int {
	return count(1, n)
}

// 计算闭区间 [low, high] 的整数能组成的 BST 数量.
func count(low, high int) int {
	// 空区间时也是一种情况, 只不过节点是 nil 而已
	if low > high {
		return 1
	}

	res := 0

	// i 作为根节点, 递归计算每个数字作为根节点时, 它左右区间的 BST 数量
	for i := low; i <= high; i++ {
		left := count(low, i-1)
		right := count(i+1, high)
		res += left * right
	}

	return res
}
