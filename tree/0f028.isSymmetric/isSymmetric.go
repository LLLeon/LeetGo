package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return recursive(root.Left, root.Right)
}

func recursive(left, right *TreeNode) bool {
	// 左右子节点同时为 nil 才是对称的
	if left == nil && right == nil {
		return true
	}

	// 非对称的情况
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}

	// 递归比较这两对节点是否对称
	return recursive(left.Left, right.Right) && recursive(left.Right, right.Left)
}
