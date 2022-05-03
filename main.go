package main

import (
	"DataStructures/binaryTree"
	"fmt"
)

func main() {
	// initialize the tree
	tree := &binaryTree.Node{
		Key: 6,
		Left: &binaryTree.Node{
			Key: 5,
			Left: &binaryTree.Node{
				Key: 4,
			},
			Right: &binaryTree.Node{
				Key: 8,
			},
		},
		Right: &binaryTree.Node{
			Key: 8,
			Left: &binaryTree.Node{
				Key: 4,
			},
			Right: &binaryTree.Node{
				Key: 7,
			},
		},
	}
	node := binaryTree.Search(tree, 5)
	fmt.Printf("%d", node.Key)
}
