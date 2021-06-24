package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	diff := math.Abs(float64(left - right))

	// 注意还要判断根节点的左右子树是不是平衡的
	return diff <= 1 && IsBalanced(root.Left) && IsBalanced(root.Right)
}

func maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := maxDepth(node.Left)
	right := maxDepth(node.Right)

	return int(math.Max(float64(left), float64(right))) + 1
}
