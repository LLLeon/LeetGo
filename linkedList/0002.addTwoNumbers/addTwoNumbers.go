package main

import "fmt"

// ListNode represents a node of a singly linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewListNode creates a new node with val.
func NewListNode(val int) *ListNode {
	return &ListNode{
		Val: val,
	}
}

// NextNode returns the node after the ln.
func (ln *ListNode) NextNode() *ListNode {
	return ln.Next
}

// InsertAfter inserts the node with val after ln and return val.
func (ln *ListNode) InsertAfter(val int) int {
	return insert(val, ln)
}

func insert(val int, at *ListNode) int {
	if at.Next == nil {
		next := &ListNode{}
		next.Val = val
		at.Next = next
	} else {
		oldNext := at.Next
		newNext := &ListNode{}

		newNext.Val = val
		newNext.Next = oldNext

		at.Next = newNext
	}

	return val
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	var (
		x int
		y int

		// 各位相加后的进位值，默认值为 0。因为各位都是 0-9 的数，
		// 考虑到进位，最大值是 9+9+1，所以 carry 不是 0 就是 1
		carry int
	)
	// 用来移动 l1 和 l2 节点
	p := l1
	q := l2

	// dummyNode 用来记录计算结果
	// curr 表示计算结果的当前节点
	dummyNode := NewListNode(0)
	curr := dummyNode

	for p != nil || q != nil {
		// 读取 l1 和 l2 当前节点的值
		if p != nil {
			x = p.Val
		} else {
			x = 0
		}

		if q != nil {
			y = q.Val
		} else {
			y = 0
		}

		// 计算当前位的和 sum，并更新进位值 carry
		sum := x + y + carry
		carry = sum / 10

		// 记录当前位计算结果
		curr.Next = NewListNode(sum % 10)
		curr = curr.Next

		// 移动到 l1 和 l2 下一个节点
		if p != nil {
			p = p.Next
		}

		if q != nil {
			q = q.Next
		}
	}

	// l1 和 l2 所有节点计算完毕后，查看进位值并记录
	if carry > 0 {
		curr.Next = NewListNode(carry)
	}

	return dummyNode.Next
}

func main() {
	l1n1 := NewListNode(2)

	l1n1.InsertAfter(4)
	l1n2 := l1n1.NextNode()

	l1n2.InsertAfter(3)
	l1n3 := l1n2.NextNode()

	fmt.Println(l1n1.Val, l1n2.Val, l1n3.Val)

	l2n1 := NewListNode(5)

	l2n1.InsertAfter(6)
	l2n2 := l2n1.NextNode()

	l2n2.InsertAfter(4)
	l2n3 := l2n2.NextNode()

	fmt.Println(l2n1.Val, l2n2.Val, l2n3.Val)

	d1 := addTwoNumbers(l1n1, l2n1)
	d2 := d1.NextNode()
	d3 := d2.NextNode()

	fmt.Println(d1.Val, d2.Val, d3.Val)
}
