package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	InvertTree(root.Left)
	InvertTree(root.Right)

	return root
}

// 利用栈来暂存节点的子节点后, 交换子节点位置.
func InvertTreeII(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	stack := []*TreeNode{root}

	for len(stack) != 0 {
		// 取出栈顶节点
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 将该节点的左右子节点入栈
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		// 交换该节点的左右子节点
		node.Left, node.Right = node.Right, node.Left
	}

	return root
}
