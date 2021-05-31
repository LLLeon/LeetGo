package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int

	left := InorderTraversal(root.Left)
	right := InorderTraversal(root.Right)

	res = append(res, left...)
	res = append(res, root.Val)
	res = append(res, right...)

	return res
}
