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

// 时间复杂度 O(logN*logN).
func CountNodesIII(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left, right := root, root
	leftHeight, rightHeight := 0, 0

	// 这两个 for 循环的时间复杂度是 O(logN).
	// 统计左子树高度
	for left != nil {
		left = left.Left
		leftHeight++
	}

	// 统计右子树高度
	for right != nil {
		right = right.Right
		rightHeight++
	}

	// 如果左右子树高度相同则这棵树是满二叉树, 它的节点数量是 2 的 H 次方.
	if leftHeight == rightHeight {
		return int(math.Pow(float64(2), float64(leftHeight)) - 1)
	}

	// 如果不是满二叉树, 则用 DFS 递归统计节点数量.
	// 到了这一步说明这是棵完全二叉树, 它的子树至少有一棵是满二叉树, 会触发 leftHeight == rightHeight 的逻辑,
	// 因此这两个递归调用, 只有一个会递归下去, 这就降低了一定的时间复杂度 (O(logN)).
	return 1 + CountNodesIII(root.Left) + CountNodesIII(root.Right)
}
