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
	max := findMaxItem(nums)
	left, right := findLeftAndRightPart(nums, max)
	root := constructNode(max)

	// 对左右部分数组进行递归处理
	root.Left = ConstructMaximumBinaryTree(left)
	root.Right = ConstructMaximumBinaryTree(right)

	return root
}

func findMaxItem(nums []int) int {
	// 注意这里, 由于题目要求数组中的数字是 [0, 1000] 中的值, 所以可以设置为 0
	max := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func findLeftAndRightPart(nums []int, max int) ([]int, []int) {
	if len(nums) == 0 {
		return []int{}, []int{}
	}

	valIndexMap := make(map[int]int)
	for i, v := range nums {
		valIndexMap[v] = i
	}

	idx, ok := valIndexMap[max]
	if !ok {
		return []int{}, []int{}
	}

	return nums[:idx], nums[idx+1:]
}

func constructNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}
