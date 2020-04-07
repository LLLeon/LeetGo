/*************************************************************************
+Author   : ChenhuiJia
+Data     : 2020-04-07
+************************************************************************/

package main

import "testing"

func TestMergeTwoLists(t *testing.T) {
	l1 := NewListNode(1)
	l1.Next = NewListNode(3)
	l1.Next.Next = NewListNode(5)

	l2 := NewListNode(2)
	l2.Next = NewListNode(4)
	l2.Next.Next = NewListNode(6)

	ml := MergeTwoLists(l1, l2)

	curr := ml
	for curr != nil {
		t.Log(curr.Val)
		curr = curr.Next
	}
}
