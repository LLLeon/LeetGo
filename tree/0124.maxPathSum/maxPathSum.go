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

		innerMaxSum := lSum + root.Val + rSum // 这是当前递归到的节点与其左右节点各自路径和的和, 也就是当前节点的最终路径和
		maxSum = max(maxSum, innerMaxSum)

		// 递归函数返回的是 (root + 左子树或右子树路径和) 或 0, 而左子树或右子树的路径和也是这么递归得来的,
		// 所以如此回溯上去, 每次递归都记录当前 root 节点和其左右节点各自路径和的和 innerMaxSum, 并将其与 maxSum 的较大者作为当前节点的最终路径和,
		// 最终递归完毕, 得到的就是要求的最大路径和
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
