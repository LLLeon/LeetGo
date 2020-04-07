/*************************************************************************
+Author   : ChenhuiJia
+Data     : 2020-04-07
+************************************************************************/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

// 类似归并排序的合并逻辑.
func MergeTwoLists(l1, l2 *ListNode) *ListNode {
	newList := NewListNode(0)
	curr := newList

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			curr = curr.Next
			l1 = l1.Next
		} else {
			curr.Next = l2
			curr = curr.Next
			l2 = l2.Next
		}
	}

	if l1 == nil {
		curr.Next = l2
	} else {
		curr.Next = l1
	}

	return newList.Next
}
