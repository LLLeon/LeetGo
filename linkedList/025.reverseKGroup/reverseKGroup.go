package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 以 k 为单位反转链表的一部分节点, 分解为子问题就是反转 N 条长度为 k 的链表, 然后把这些反转后的链表连接起来.
func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	left := head
	right := head

	// 将 right 移动到下一个要反转的子链表, 将其作为子链表的 head
	for i := 0; i < k; i++ {
		// 如果不足 k 个节点则不反转, 直接返回 head 节点
		if right == nil {
			return head
		}

		right = right.Next
	}

	// 反转子链表
	newHead := reverse(left, right)

	// 连接反转后的子链表
	left.Next = ReverseKGroup(right, k)

	// 返回反转后的子链表的 head
	return newHead
}

// 反转 [left, right) 区间的节点
func reverse(left, right *ListNode) *ListNode {
	if left == nil {
		return nil
	}

	var (
		pre  *ListNode
		curr = left
	)

	for {
		// 和反转一整条链表的区别就是这里
		if curr == right {
			break
		}

		temp := curr
		curr = curr.Next
		temp.Next = pre
		pre = temp
	}

	return pre
}
