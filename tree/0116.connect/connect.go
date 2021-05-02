package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func Connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	connectTwoNodes(root.Left, root.Right)
	return root
}

func connectTwoNodes(node1, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}

	// 连接传入的两个节点
	node1.Next = node2

	// 连接传入节点的左右子节点
	connectTwoNodes(node1.Left, node1.Right)
	connectTwoNodes(node2.Left, node2.Right)

	// 连接非同一父节点的两个相邻节点
	connectTwoNodes(node1.Right, node2.Left)
}
