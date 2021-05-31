package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lDepth := MaxDepth(root.Left)
	rDepth := MaxDepth(root.Right)
	depth := max(lDepth, rDepth) + 1

	return depth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
