package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 时间复杂度 O(N), 空间复杂度 O(N).
func KthSmallest(root *TreeNode, k int) int {
	res := make([]int, 0)
	traverse(root, &res)
	return res[k-1]
}

func traverse(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	traverse(root.Left, res)
	*res = append(*res, root.Val)
	traverse(root.Right, res)
}

// 时间复杂度 O(N), 空间复杂度 O(1).
func KthSmallestII(root *TreeNode, k int) int {
	res, rank := 0, 0
	inorder(root, k, &rank, &res)

	return res
}

func inorder(root *TreeNode, k int, rank *int, res *int) {
	if root == nil {
		return
	}

	inorder(root.Left, k, rank, res)
	*rank++
	if *rank == k {
		*res = root.Val
		return
	}
	inorder(root.Right, k, rank, res)
}
