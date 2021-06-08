package main

import (
	"math"
	"sync"
)

const (
	// linkedHead and linkedTail represents the head
	// and tail of a bidirectional linked list without
	// occupying the linked list space.
	linkedHead = math.MaxUint16
	linkedTail = math.MaxUint16
)

// LinkedNode represents a node in a bidirectional linked list.
type LinkedNode struct {
	key   int
	value int
	pre   *LinkedNode
	next  *LinkedNode
}

// NewLinkedNode returns a new LinkedNode.
func NewLinkedNode(key, value int) *LinkedNode {
	return &LinkedNode{
		key:   key,
		value: value,
	}
}

// LRUCache represents a LRU cache implemented with hash
// tables and bidirectional linked list.
type LRUCache struct {
	count int
	size  int
	cache map[int]*LinkedNode
	lock  sync.Mutex
}

// Constructor returns a LRUCache with a specified capacity.
func Constructor(capacity int) *LRUCache {
	if capacity <= 0 {
		return nil
	}

	head := NewLinkedNode(linkedHead, -1)
	head.pre = nil

	tail := NewLinkedNode(linkedTail, -1)
	tail.next = nil

	head.next = tail
	tail.pre = head

	lc := &LRUCache{
		count: 0,
		size:  capacity,
		cache: make(map[int]*LinkedNode),
	}

	lc.cache[linkedHead] = head
	lc.cache[linkedTail] = tail

	return lc
}

func (lc *LRUCache) Put(key, value int) {
	lc.lock.Lock()
	defer lc.lock.Unlock()

	if key == linkedHead || key == linkedTail {
		return
	}

	var (
		newHead  *LinkedNode
		exists   bool
		isUpdate bool
		lhead    = lc.cache[linkedHead]
		ltail    = lc.cache[linkedTail]
	)

	// If the key you want to set exists, remove it from its
	// original location and put it in the head of the linked list.
	newHead, exists = lc.cache[key]
	if exists {
		originalPre := newHead.pre
		originalPre.next = newHead.next
		newHead.next.pre = originalPre

		// Temporarily subtract 1 number of nodes.
		lc.count--

		// update original value
		newHead.value = value
		isUpdate = true
	} else {
		newHead = NewLinkedNode(key, value)
	}

	// Special handling is required when there are no
	// nodes in the linked list.
	if lc.count == 0 {
		lhead.next = newHead
		newHead.pre = lhead
		newHead.next = ltail
		ltail.pre = newHead

		lc.cache[key] = newHead
		lc.count++
		return
	}

	// When the cache is full and not updating old values,
	// delete the tail node and put the new node in the head.
	if lc.count == lc.size && !isUpdate {
		nodeRemove := ltail.pre
		nodeRemove.pre.next = ltail
		ltail.pre = nodeRemove.pre

		delete(lc.cache, nodeRemove.key)
		nodeRemove = nil
		lc.count--
	}

	// Insert a new node in the head of the linked list.
	oldHead := lhead.next
	lhead.next = newHead
	newHead.pre = lhead
	newHead.next = oldHead
	oldHead.pre = newHead

	// Don't forget to cache the node.
	lc.cache[key] = newHead
	lc.count++
}

// Get returns the specified value in a bidirectional linked list.
func (lc *LRUCache) Get(key int) int {
	lc.lock.Lock()
	defer lc.lock.Unlock()

	if lc.count == 0 || key == linkedHead || key == linkedTail {
		return -1
	}

	dstNode, ok := lc.cache[key]
	if !ok {
		return -1
	}

	v := dstNode.value
	lhead := lc.cache[linkedHead]

	// If the target value is in the head of the list,
	// return it directly.
	if dstNode.pre == lhead {
		return v
	}

	// Delete the target node from its original location.
	preNode := dstNode.pre
	preNode.next = dstNode.next
	dstNode.next.pre = preNode

	// Move the target node to the head of the linked list.
	oldHead := lhead.next
	lhead.next = dstNode
	dstNode.pre = lhead
	dstNode.next = oldHead
	oldHead.pre = dstNode

	return v
}
