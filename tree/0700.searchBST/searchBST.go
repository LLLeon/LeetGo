package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 时间复杂度 O(H), H 为树高, 平均时间复杂度为 O(logN), 最坏时间复杂度为 O(N).
// 空间复杂度 O(H), 递归栈的深度为 H, 平均情况下深度为 O(logN), 最坏情况下深度为 O(N).
func SearchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}

	if val < root.Val {
		return SearchBST(root.Left, val)
	}

	if val > root.Val {
		return SearchBST(root.Right, val)
	}

	return nil
}

// 时间复杂度同上.
// 空间复杂度 O(1).
func SearchBSTII(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
		if val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	// 此时要么找到目标节点, 要么为 nil.
	return root
}
