package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GenerateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	return build(1, n)
}

func build(low, high int) []*TreeNode {
	var res = make([]*TreeNode, 0)

	if low > high {
		res = append(res, nil)
		return res
	}

	// 为每个节点 i 构造 BST
	for i := low; i <= high; i++ {
		// 递归的构造节点 i 的左右子树
		leftTree := build(low, i-1)
		rightTree := build(i+1, high)

		// 为节点 i 组合每一种左右子树的可能
		for _, left := range leftTree {
			for _, right := range rightTree {
				root := &TreeNode{Val: i}
				root.Left = left
				root.Right = right
				// 将节点 i 添加到结果
				res = append(res, root)
			}
		}
	}

	return res
}
