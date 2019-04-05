package main

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	s := Constructor()

	s.Push(4)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(1)
	s.Push(6)

	t.Logf("data stack: %v\n", s.dataStack)
	t.Logf("min index stack: %v\n", s.minIndex)

	min, _ := s.GetMin()
	t.Logf("min: %d\n", min)

	s.Pop()
	t.Logf("data stack: %v\n", s.dataStack)
	t.Logf("min index stack: %v\n", s.minIndex)

	s.Pop()
	t.Logf("data stack: %v\n", s.dataStack)
	t.Logf("min index stack: %v\n", s.minIndex)

	min, _ = s.GetMin()
	t.Logf("min: %d\n", min)

	s.Pop()
	t.Logf("data stack: %v\n", s.dataStack)
	t.Logf("min index stack: %v\n", s.minIndex)

	min, _ = s.GetMin()
	t.Logf("min: %d\n", min)
}
