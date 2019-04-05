package cycle

// ListNode represents a node of a linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewListNode creates a new node with val.
func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func HasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head.Next

	for slow != nil && fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}
