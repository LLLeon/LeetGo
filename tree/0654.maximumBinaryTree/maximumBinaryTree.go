package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ConstructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return constructNode(nums[0])
	}

	// 找到数组中的最大值和其左右部分数组
	max, idx := findMaxItem(nums)
	left := nums[:idx]
	right := nums[idx+1:]
	root := constructNode(max)

	// 对左右部分数组进行递归处理
	root.Left = ConstructMaximumBinaryTree(left)
	root.Right = ConstructMaximumBinaryTree(right)

	return root
}

func findMaxItem(nums []int) (max int, index int) {
	// 注意这里, 由于题目要求数组中的数字是 [0, 1000] 中的值, 所以可以设置为 0
	for i, v := range nums {
		if v > max {
			max = v
			index = i
		}
	}
	return max, index
}

func constructNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}
