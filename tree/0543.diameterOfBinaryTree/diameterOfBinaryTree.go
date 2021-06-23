package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func DiameterOfBinaryTree(root *TreeNode) int {
	res = 1
	recursive(root)

	return res - 1
}

func recursive(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := recursive(node.Left)
	right := recursive(node.Right)

	res = max(res, left+right+1)
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
