package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ------------------------------------------- preorder ----------------------------------------------

// 记录 BST 的最大节点和
var maxSum int

// 很奇怪, 在 LeetCode 单独测试 [4,3,null,1,2] 这个输入时结果是正确的, 但提交代码后却报错, 与单独测试的结果不一样.
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
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ----------------------------------------- postorder -------------------------------------------------

const (
	isBST  = 1
	notBST = 0
)

func MaxSumBSTII(root *TreeNode) int {
	traverseII(root)
	return maxSum
}

// 返回的数组存储四个值, 分别代表 [isBST, min, max, sum]
// isBST: 以 root 为根节点的二叉树是不是 BST;
// min 和 max: 以 root 为根节点的二叉树中所有节点的最小值和最大值;
// sum: 以 root 为根节点的二叉树中所有节点的和.
func traverseII(root *TreeNode) [4]int {
	if root == nil {
		// 注意这个情况是 BST
		return [4]int{isBST, math.MaxInt64, math.MinInt64, 0}
	}

	left := traverseII(root.Left)
	right := traverseII(root.Right)

	res := [4]int{}

	// 判断以 root 节点为根节点的二叉树是不是 BST, 其中 root.Val 需要与左子树中的最大值和右子树中的最小值比较
	if left[0] == isBST && right[0] == isBST && root.Val > left[2] && root.Val < right[1] {
		res[0] = isBST

		// 记录以 root 节点为根节点的 BST 中的最小值和最大值
		res[1] = min(left[1], root.Val)
		res[2] = max(right[2], root.Val)

		// 计算以 root 节点为根节点的 BST 的所有节点和
		res[3] = left[3] + root.Val + right[3]

		// 记录最大节点和
		maxSum = max(maxSum, res[3])
	} else {
		res[0] = notBST
	}

	return res
}
