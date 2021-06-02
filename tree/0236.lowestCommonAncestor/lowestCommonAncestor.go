package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// base case
	if root == nil {
		return nil
	}
	// 一个节点也是其自身的祖先.
	// 不管另一个节点在不在树中, 在树中的节点是它们两个的祖先.
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	leftLCA := LowestCommonAncestor(root.Left, p, q)
	rightLCA := LowestCommonAncestor(root.Right, p, q)

	// 1. 这种情况, p 和 q 不可能既同时在 root 的左子树中又同时在 root 的右子树中有 LCA,
	// 根据 base case, 只能是 p 或 q 分别就是 leftLCA 或 rightLCA, 那么它们的 LCA 就是 root.
	if leftLCA != nil && rightLCA != nil {
		return root
	}

	// 2. 在 root 的子树中都没有找到 p 和 q 的 LCA, 返回 nil.
	if leftLCA == nil && rightLCA == nil {
		return nil
	}

	// 3. 只在左/右子树中找到 LCA, 返回它.
	if leftLCA == nil {
		return rightLCA
	}

	return leftLCA
}
