package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ConvertBST(root *TreeNode) *TreeNode {
	sum := 0
	traverse(root, &sum)

	return root
}

func traverse(root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	// 依然是中序遍历, 不过是先遍历右子树, 降序遍历, 这样就可以保证必然是从大到小遍历节点, 累加和即可
	traverse(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	traverse(root.Left, sum)
}
