package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func KthSmallest(root *TreeNode, k int) int {
	res := make([]int, 0)
	traverse(root, &res)
	return res[k-1]
}

func traverse(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	traverse(root.Left, res)
	*res = append(*res, root.Val)
	traverse(root.Right, res)
}
