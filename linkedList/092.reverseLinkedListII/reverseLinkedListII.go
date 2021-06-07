package main

var successor *ListNode

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
		successor = head.Next
		return head
	}

	// 以 head.Next 为起点, 需要反转前 n - 1 个节点
	last := reverseN(head.Next, n-1)

	// 反转节点指向
	head.Next.Next = head
	// 连接前面记录的后继节点
	head.Next = successor

	return last
}

func ReverseBetweenII(head *ListNode, left, right int) *ListNode {
	if head == nil || left > right {
		return nil
	}

	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head

	// 找到待反转区域中第一个节点的前一个节点
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	// curr 表示待反转区域的第一个节点
	curr := pre.Next
	for i := 0; i < right-left; i++ {
		next := curr.Next     // next 表示 curr 的下一个节点
		curr.Next = next.Next // 把 curr 的下一个节点指向 next 的下一个节点
		next.Next = pre.Next  // 把 next 的下一个节点指向 pre 的下一个节点
		pre.Next = next       // 把 pre 的下一个节点指向 next
	}

	return dummyNode.Next
}
