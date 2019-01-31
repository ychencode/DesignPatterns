package main

import "fmt"

func main() {

	avlTree := NewAvlTree()

	for i := 0; i < 100; i++ {
		avlTree.root = InsertTreeNode(avlTree.root, i)
	}

	preOrderVisit(avlTree.root)

	fmt.Println(FindNode(avlTree.root, 3))

}

func preOrderVisit(node *TreeNode) {
	if node != nil {
		fmt.Println(node.key)
		preOrderVisit(node.left)
		preOrderVisit(node.right)
	}
}
