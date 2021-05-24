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

// 时间复杂度 O(N2), 空间复杂度 O(N).
func NumTreesII(n int) int {
	dp := make([]int, n+1)
	// 两种特殊情况
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		for j := 0; j <= i-1; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}

	return dp[n]
}
