package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// # 代表空节点.
func (c *Codec) Serialize(root *TreeNode) string {
	if root == nil {
		return "#,"
	}

	left := c.Serialize(root.Left)
	right := c.Serialize(root.Right)

	// preorder
	return strconv.Itoa(root.Val) + "," + left + right
}

func (c *Codec) Deserialize(data string) *TreeNode {
	nodes := strings.Split(data, ",")
	return c.buildTree(&nodes)
}

func (c *Codec) buildTree(nodes *[]string) *TreeNode {
	nodeVal := "#"

	if len(*nodes) > 0 {
		nodeVal = (*nodes)[0]
		*nodes = (*nodes)[1:]
	}

	if nodeVal == "#" {
		return nil
	}

	val, _ := strconv.Atoi(nodeVal)
	root := &TreeNode{Val: val}
	root.Left = c.buildTree(nodes)
	root.Right = c.buildTree(nodes)

	return root
}
