package main

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func FindDuplicateSubtrees(root *TreeNode) []*TreeNode {
	counter := make(map[string]int)
	res := make([]*TreeNode, 0)

	traverse(root, counter, &res)
	return res
}

func traverse(root *TreeNode, counter map[string]int, res *[]*TreeNode) string {
	if root == nil {
		return "#"
	}

	// 构造该 root 节点的子树
	left := traverse(root.Left, counter, res)
	right := traverse(root.Right, counter, res)
	subTree := strconv.Itoa(root.Val) + "," + left + "," + right

	// 记录子树出现的次数, 如果有相同的就将其 root 节点计入结果
	counter[subTree]++
	if counter[subTree] == 2 {
		*res = append(*res, root)
	}

	return subTree
}
