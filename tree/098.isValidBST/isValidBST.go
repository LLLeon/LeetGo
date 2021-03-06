package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 时间复杂度 O(N), 空间复杂度 O(N), N 为二叉树节点数量.
func IsValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}

	// 对左右子树的判断, 如果某个节点值 root.Val 不在合法范围内, 则表示这棵二叉树不是 BST,
	// 只不过左右子树的最大/小值范围不同.
	if root.Val <= lower || root.Val >= upper {
		return false
	}

	// 限定左子树的最大值为 root.Val, 右子树的最小值为 root.Val, 递归判断
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

// 时间复杂度 O(N), 空间复杂度 O(N), N 为二叉树节点数量.
func IsValidBSTII(root *TreeNode) bool {
	prev := math.MinInt64
	return traverse(root, &prev)
}

func traverse(root *TreeNode, prev *int) bool {
	if root == nil {
		return true
	}

	// 如果左子树不满足条件, 直接返回 false
	if !traverse(root.Left, prev) {
		return false
	}
	// 如果当前节点小于它的左节点, 直接返回 false
	if root.Val <= *prev {
		return false
	}
	// 如果满足条件, 将当前节点作为 prev 继续遍历
	*prev = root.Val

	// 同样判断右子树
	return traverse(root.Right, prev)
}
