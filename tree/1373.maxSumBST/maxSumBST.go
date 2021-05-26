package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 记录 BST 的最大节点和
var maxSum int

func MaxSumBST(root *TreeNode) int {
	traverse(root)
	return maxSum
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	var (
		leftMax  int
		rightMin int
		leftSum  int
		rightSum int
		rootSum  int
	)

	// 判断当前节点的左右子树是不是 BST
	if !isValidBST(root.Left) || !isValidBST(root.Right) {
		goto next
	}

	// 判断以当前节点为 root 的二叉树是不是 BST
	leftMax = findMaxVal(root.Left)
	rightMin = findMinVal(root.Right)
	if root.Val <= leftMax || root.Val >= rightMin {
		goto next
	}

	// 计算以当前节点为 root 的 BST 的节点和
	leftSum = getSum(root.Left)
	rightSum = getSum(root.Right)
	rootSum = leftSum + root.Val + rightSum

	// 记录最大的 BST 节点和
	maxSum = max(maxSum, rootSum)

next:
	// 递归计算左右子树
	traverse(root.Left)
	traverse(root.Right)
}

func isValidBST(root *TreeNode) bool {
	prev := math.MinInt64
	return isValidBSTHelper(root, &prev)
}

func isValidBSTHelper(root *TreeNode, prev *int) bool {
	if root == nil {
		return true
	}

	if !isValidBSTHelper(root.Left, prev) {
		return false
	}

	if root.Val <= *prev {
		return false
	}
	*prev = root.Val

	return isValidBSTHelper(root.Right, prev)
}

func findMaxVal(root *TreeNode) int {
	if root == nil {
		return math.MinInt64
	}
	for root.Right != nil {
		root = root.Right
	}
	return root.Val
}

func findMinVal(root *TreeNode) int {
	if root == nil {
		return math.MaxInt64
	}
	for root.Left != nil {
		root = root.Left
	}
	return root.Val
}

func getSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := getSum(root.Left)
	right := getSum(root.Right)
	sum := left + right + root.Val

	return sum
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
