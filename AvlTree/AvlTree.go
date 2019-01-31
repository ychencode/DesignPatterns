package main

type AvlTree struct {
	root *TreeNode
}

func NewAvlTree() *AvlTree {
	return new(AvlTree)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(height(node.left), height(node.right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getBalance(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return height(node.left) - height(node.right)
}

func leftRotate(pivot *TreeNode) *TreeNode {
	yNode := pivot.right
	tNode := yNode.left

	yNode.left = pivot
	pivot.right = tNode

	pivot.height = max(height(pivot.left), height(pivot.right)) + 1
	yNode.height = max(height(yNode.left), height(yNode.right)) + 1

	return yNode
}

func rightRotate(pivot *TreeNode) *TreeNode {
	xNode := pivot.left
	tNode := xNode.right

	xNode.right = pivot
	pivot.left = tNode

	xNode.height = max(height(xNode.left), height(xNode.right)) + 1
	pivot.height = max(height(pivot.left), height(pivot.right)) + 1
	return xNode
}

func InsertTreeNode(node *TreeNode, key int) *TreeNode {
	if node == nil {
		return NewTreeNode(key)
	}
	if key < node.key {
		node.left = InsertTreeNode(node.left, key)
	} else if key > node.key {
		node.right = InsertTreeNode(node.right, key)
	} else {
		return node
	}

	node.height = max(height(node.left), height(node.right)) + 1

	balance := getBalance(node)

	if balance > 1 && key < node.left.key {
		// left left
		return rightRotate(node)
	} else if balance > 1 && key > node.left.key {
		// left right
		node.left = leftRotate(node.left)
		return rightRotate(node)
	} else if balance < -1 && key < node.right.key {
		// right left
		node.right = rightRotate(node.right)
		return leftRotate(node)
	} else if balance < -1 && key > node.right.key {
		// right right
		return leftRotate(node)
	}

	return node
}

func FindNode(node *TreeNode, key int) bool {
	if node == nil {
		return false
	}
	if node.key == key {
		return true
	} else if key < node.key {
		return FindNode(node.left, key)
	} else {
		return FindNode(node.right, key)
	}
}
