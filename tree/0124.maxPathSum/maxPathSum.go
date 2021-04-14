package maxPathSum

import (
	"math"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func MaxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var dfs func(root *TreeNode) int

	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		lSum := dfs(root.Left)
		rSum := dfs(root.Right)

		innerMaxSum := lSum + root.Val + rSum
		maxSum = max(maxSum, innerMaxSum)

		// 返回的是 (root + 左子树或右子树的和) 或 0, 而左子树或右子树的值也是这么递归得来的,
		// 所以如此回溯上去, 每次递归都记录当前 root 节点和其左右节点各自路径和的和, 并将其与 maxSum 的较大者作为最终路径和
		res := root.Val + max(lSum, rSum)
		return max(0, res)
	}

	dfs(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
