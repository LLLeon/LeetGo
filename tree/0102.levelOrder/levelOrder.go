package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	q := []*TreeNode{root}
	res := [][]int{}

	for len(q) != 0 {
		innerRes := []int{}
		levelQ := []*TreeNode{}

		for _, node := range q {
			innerRes = append(innerRes, node.Val)
			if node.Left != nil {
				levelQ = append(levelQ, node.Left)
			}
			if node.Right != nil {
				levelQ = append(levelQ, node.Right)
			}
		}

		q = levelQ[:]
		res = append(res, innerRes)
	}

	return res
}
