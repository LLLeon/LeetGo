package traversal

type Node struct {
	Val      int
	Children []*Node
}

func LevelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := make([][]int, 0)
	// 把 root 节点入队
	q := []*Node{root}

	// 开始遍历队列
	for len(q) != 0 {
		innerRes := make([]int, 0)
		innerQ := make([]*Node, 0)

		// 取出队列中的节点, 并将其子节点入队
		for _, node := range q {
			innerRes = append(innerRes, node.Val)
			innerQ = append(innerQ, node.Children...)
		}

		// 将当前层的节点值添加到结果数组
		res = append(res, innerRes)
		q = innerQ[:]
	}

	return res
}
