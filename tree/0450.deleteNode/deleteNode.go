package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DeleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	// 找到 key
	if key == root.Val {
		// 1) root 的左右节点都是 nil
		if root.Left == nil && root.Right == nil {
			return nil
		}
		// 2) root 的左节点是 nil, 用其右节点替换到自己的位置
		if root.Left == nil {
			return root.Right
		}
		// 3) root 的右节点是 nil, 用其左节点替换到自己的位置
		if root.Right == nil {
			return root.Left
		}
		// 4) root 的左右节点都不为 nil, 用其左子树中的最大节点 (Predecessor: 前驱节点)
		// 或右子树中的最小节点 (Successor: 后继节点) 替换自己的位置
		if root.Left != nil && root.Right != nil {
			maxNode := getMaxNode(root.Left)
			root.Val = maxNode.Val
			// 删除左子树中的最大节点
			root.Left = DeleteNode(root.Left, maxNode.Val)
		}
	}

	if key < root.Val {
		root.Left = DeleteNode(root.Left, key)
	}

	if key > root.Val {
		root.Right = DeleteNode(root.Right, key)
	}

	return root
}

func getMaxNode(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	for root.Right != nil {
		root = root.Right
	}

	return root
}
