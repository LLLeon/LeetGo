package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InsertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val < root.Val {
		// 此时 val 小于 root 的值, 将其插入 root 的左子树,
		// 而这里 InsertIntoBST() 返回的是将 val 插入 root 的左子树之后返回的根节点,
		// 将其作为 root 的左节点.
		root.Left = InsertIntoBST(root.Left, val)
	}

	if val > root.Val {
		root.Right = InsertIntoBST(root.Right, val)
	}

	return root
}
