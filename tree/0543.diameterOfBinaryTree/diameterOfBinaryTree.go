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

	res = max(res, left+right+1) // 更新 res
	return max(left, right) + 1  // 以当前节点为根节点的最大深度, 这个返回值会作为二叉树直径的一部分(即 left 或 right 的值)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
