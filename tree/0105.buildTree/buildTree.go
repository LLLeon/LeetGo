package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(preorder, inorder []int) *TreeNode {
	if len(preorder) != len(inorder) {
		return nil
	}

	// 空间换时间, 先把中序遍历序列的 [value: index] 键值对存起来,
	// 以便后面可以找到前序遍历序列中的 root 节点在中序遍历序列中的位置
	valIdxMap := make(map[int]int, len(inorder))

	for i, v := range inorder {
		valIdxMap[v] = i
	}

	return build(preorder, 0, len(preorder)-1, valIdxMap, 0, len(inorder)-1)
}

func build(preorder []int, preStart, preEnd int, valIdxMap map[int]int, inStart, inEnd int) *TreeNode {
	if preStart > preEnd || inStart > inEnd {
		return nil
	}

	// 找到 root 节点及其在 inorder 中的位置
	rootVal := preorder[preStart]
	root := &TreeNode{Val: rootVal}
	pivotIdx := valIdxMap[rootVal]

	// 根据 root 节点在 inorder 中的位置确定左子树的界限
	preLeftStart := preStart + 1
	preLeftEnd := pivotIdx - inStart + preStart
	inLeftStart := inStart
	inLeftEnd := pivotIdx - 1

	// 根据 root 节点在 inorder 中的位置确定右子树的界限
	preRightStart := pivotIdx - inStart + preStart + 1
	preRightEnd := preEnd
	inRightStart := pivotIdx + 1
	inRightEnd := inEnd

	root.Left = build(preorder, preLeftStart, preLeftEnd, valIdxMap, inLeftStart, inLeftEnd)
	root.Right = build(preorder, preRightStart, preRightEnd, valIdxMap, inRightStart, inRightEnd)

	return root
}
