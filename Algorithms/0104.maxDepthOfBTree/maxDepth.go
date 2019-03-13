package maxDepthOfBTree

import (
	"math"
)

// Node represents a node in a binary tree.
type Node struct {
	data  int
	left  *Node
	right *Node
}

// MaxDepth returns the max depth of the binary tree.
func MaxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	lDepth := MaxDepth(root.left)
	rDepth := MaxDepth(root.right)
	max := int(math.Max(float64(lDepth), float64(rDepth))) + 1

	return max
}
