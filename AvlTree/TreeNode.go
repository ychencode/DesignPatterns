package main

type TreeNode struct {
	key    int
	height int
	left   *TreeNode
	right  *TreeNode
}

func NewTreeNode(key int) *TreeNode {
	node := &TreeNode{
		key:    key,
		height: 1,
		left:   nil,
		right:  nil,
	}
	return node
}


