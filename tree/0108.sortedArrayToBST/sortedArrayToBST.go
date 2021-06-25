package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	// 选取数组中间位置的数字作为根节点
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	// 递归构建根节点的左右子节点
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)

	return root
}
