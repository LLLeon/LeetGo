package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Flatten(root *TreeNode) {
	if root == nil {
		return
	}

	// 记录当前节点的左右子树
	left := root.Left
	right := root.Right

	// 把左子树作为右子树
	root.Left = nil
	root.Right = left

	// 然后遍历到现在右子树的最下层节点, 把之前的右子树接到现在的右子树下
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right

	// 递归右子树 (左子树已经为 nil)
	Flatten(root.Right)
}
