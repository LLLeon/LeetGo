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
