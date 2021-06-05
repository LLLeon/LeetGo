package main

import (
	"math"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	level := 0

	for node := root; node.Left != nil; node = node.Left {
		level++
	}

	return sort.Search(1<<(level+1), func(k int) bool {
		if k <= 1<<level {
			return false
		}

		bits := 1 << (level - 1)
		node := root

		for node != nil && bits > 0 {
			if bits&k == 0 {
				node = node.Left
			} else {
				node = node.Right
			}
			bits >>= 1
		}

		return node == nil
	}) - 1
}

func CountNodesII(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := CountNodesII(root.Left)
	right := CountNodesII(root.Right)
	return left + right + 1
}

func CountNodesIII(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left, right := root, root
	leftHeight, rightHeight := 0, 0

	for left != nil {
		left = left.Left
		leftHeight++
	}

	for right != nil {
		right = right.Right
		rightHeight++
	}

	if leftHeight == rightHeight {
		return int(math.Pow(float64(2), float64(leftHeight)) - 1)
	}

	return 1 + CountNodesIII(root.Left) + CountNodesIII(root.Right)
}
