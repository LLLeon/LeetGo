package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	res        = make([]*TreeNode, 0)
	subTreeMap = make(map[string]int)
)

func FindDuplicateSubtrees(root *TreeNode) []*TreeNode {
	traverse(root)
	return res
}

func traverse(root *TreeNode) string {
	if root == nil {
		return "#"
	}

	left := traverse(root.Left)
	right := traverse(root.Right)

	subTree := left + "," + right + "," + strconv.Itoa(root.Val)

	freq := subTreeMap[subTree]
	if freq == 1 {
		res = append(res, root)
	}
	subTreeMap[subTree] = freq + 1

	return subTree
}
