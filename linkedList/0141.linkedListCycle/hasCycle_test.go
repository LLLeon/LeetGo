package cycle

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	head := NewListNode(3)
	first := NewListNode(2)
	head.Next = first

	second := NewListNode(0)
	first.Next = second

	third := NewListNode(-4)
	second.Next = third

	// cycle
	third.Next = first

	if ok := HasCycle(head); !ok {
		t.Errorf("error, want: true, get: %v\n", ok)
	} else {
		t.Log("success")
	}
}
