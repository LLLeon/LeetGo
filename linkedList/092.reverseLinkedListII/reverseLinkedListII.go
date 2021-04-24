package main

var sucessor *ListNode

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseBetween(head *ListNode, left, right int) *ListNode {
	// 如果起始位置 left 为 1, 则表示要反转前 right 个节点
	if left == 1 {
		return reverseN(head, right)
	}

	// 如果起始位置不是 1, 可以以 head.Next 作为头节点, 相应的开始位置和结束位置也各要缩小 1,
	// 这样不断递归推进到起始位置为 1 的情况, 完成反转
	head.Next = ReverseBetween(head.Next, left-1, right-1)

	return head
}

// 反转前 n 个节点
func reverseN(head *ListNode, n int) *ListNode {
	// 递归到第 n 个节点, 返回其自身
	if n == 1 {
		// 记录第 n+1 个节点
		sucessor = head.Next
		return head
	}

	// 以 head.Next 为起点, 需要反转前 n - 1 个节点
	last := reverseN(head.Next, n-1)

	// 反转节点指向
	head.Next.Next = head
	// 连接前面记录的后继节点
	head.Next = sucessor

	return last
}
