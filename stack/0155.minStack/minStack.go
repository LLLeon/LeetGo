package main

/*
 ArrayStack
*/

// ArrayStack represents a stack implemented with array.
type ArrayStack struct {
	items []int
	count int
}

// NewArrayStack returns a new ArrayStack with specified size.
func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		items: make([]int, 0),
		count: 0,
	}
}

// Push an item to the top of the stack.
func (as *ArrayStack) Push(x int) {
	as.items = append(as.items, x)
	as.count++
}

// Pop up an item from the top of the stack.
func (as *ArrayStack) Pop() (int, bool) {
	if as.count == 0 {
		return 0, false
	}

	res := as.items[as.count-1]
	as.items = as.items[:as.count-1]
	as.count--

	return res, true
}

// Top returns the item at the top of the stack.
func (as *ArrayStack) Top() (int, bool) {
	if as.count == 0 {
		return 0, false
	}

	return as.items[as.count-1], true
}

/*
 MinStack
*/

// MinStack represents a struct with two array stack.
// minIndex is used to store the index of the smallest
// item in dataStack.
type MinStack struct {
	dataStack *ArrayStack
	minIndex  *ArrayStack
}

// Constructor returns a new MinStack.
func Constructor() *MinStack {
	return &MinStack{
		dataStack: NewArrayStack(),
		minIndex:  NewArrayStack(),
	}
}

// Push an item to the top of the stack.
func (ms *MinStack) Push(x int) {
	ms.dataStack.Push(x)
	lastIndex := ms.dataStack.count - 1

	// if no index in minIndex, push lastIndex to it.
	minIndex, exist := ms.minIndex.Top()
	if !exist {
		ms.minIndex.Push(lastIndex)
		return
	}

	if x < ms.dataStack.items[minIndex] {
		ms.minIndex.Push(lastIndex)
	}
}

// Pop up an item from the top of the stack.
// If the index of the pop-up item is at the top
// of minIndex, it pops up from the top of minIndex.
func (ms *MinStack) Pop() (int, bool) {
	popIndex := ms.dataStack.count - 1
	minIndex, exist := ms.minIndex.Top()

	if exist && (popIndex == minIndex) {
		ms.minIndex.Pop()
	}

	pop, ok := ms.dataStack.Pop()
	if !ok {
		return 0, false
	}

	return pop, true
}

// Top returns the item at the top of dataStack.
func (ms *MinStack) Top() (int, bool) {
	top, exist := ms.dataStack.Top()
	if !exist {
		return 0, false
	}

	return top, true
}

// GetMin returns the smallest item in MinStack.
func (ms *MinStack) GetMin() (int, bool) {
	minIndex, exist := ms.minIndex.Top()
	if !exist {
		return 0, false
	}

	min := ms.dataStack.items[minIndex]
	return min, true
}
