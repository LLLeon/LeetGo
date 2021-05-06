package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteNode(node *ListNode) {
	if node == nil {
		return
	}

	// 把当前节点的下一个节点复制到当前节点, 然后删除之, 妙啊
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
