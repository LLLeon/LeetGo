package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(inorder, postorder []int) *TreeNode {
	if len(inorder) != len(postorder) {
		return nil
	}

	valIdxMap := make(map[int]int, len(inorder))

	for i, v := range inorder {
		valIdxMap[v] = i
	}

	return build(valIdxMap, 0, len(inorder), postorder, 0, len(postorder)-1)
}

func build(valIdxMap map[int]int, inStart, inEnd int, postorder []int, postStart, postEnd int) *TreeNode {
	if inStart > inEnd || postStart > postEnd {
		return nil
	}

	// 找到 root 节点及其在 inorder 中的位置
	rootVal := postorder[postEnd]
	root := &TreeNode{Val: rootVal}
	pivotIdx := valIdxMap[rootVal]
	leftSize := pivotIdx - inStart

	// 根据 root 节点在 inorder 中的位置确定左子树的界限
	inLeftStart := inStart
	inLeftEnd := pivotIdx - 1
	postLeftStart := postStart
	postLeftEnd := postStart + leftSize - 1

	// 根据 root 节点在 inorder 中的位置确定右子树的界限
	inRightStart := pivotIdx + 1
	inRightEnd := inEnd
	postRightStart := postStart + leftSize
	postRightEnd := postEnd - 1

	root.Left = build(valIdxMap, inLeftStart, inLeftEnd, postorder, postLeftStart, postLeftEnd)
	root.Right = build(valIdxMap, inRightStart, inRightEnd, postorder, postRightStart, postRightEnd)

	return root
}
