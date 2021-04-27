package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 普通的字符串或数组回文判断可以通过双指针法实现，但链表无法向前遍历。
// 可以先把链表存入数组，然后使用双指针法来实现。
func IsPalindromeI(head *ListNode) bool {
	if head == nil {
		return true
	}

	list := make([]*ListNode, 0)

	newHead := head
	for newHead != nil {
		list = append(list, newHead)
		newHead = newHead.Next
	}

	left := 0
	right := len(list) - 1

	for left < right {
		if list[left].Val != list[right].Val {
			return false
		}

		left++
		right--
	}

	return true
}

// 正序遍历的头节点
var left *ListNode

// 类似前序后序遍历二叉树, 也可以前序后序遍历链表.
// 利用递归反转链表的原理来倒序遍历链表. 递归遍历链表也就是把链表节点入栈再出栈.
func IsPalindromeII(head *ListNode) bool {
	left = head

	return traverse(head)
}

func traverse(right *ListNode) bool {
	// 遍历到最后一个节点的后面, 递归开始返回
	if right == nil {
		return true
	}

	res := traverse(right.Next)
	// 这个解法有可以改进的地方, 因为在遇到两个节点不同时就可以终止后面的判断了, 而这里把所有节点都遍历了一遍
	res = res && (right.Val == left.Val)
	// 正序遍历到下个节点
	left = left.Next

	return res
}
