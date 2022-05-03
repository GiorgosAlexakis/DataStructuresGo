package main

import (
	"DataStructures/binaryTree"
	"fmt"
)

// have to write tests for each data structure instead of just testing with main
func main() {
	// initialize the tree
	tree := &binaryTree.Node{
		Key: 6,
		Left: &binaryTree.Node{
			Key: 4,
			Left: &binaryTree.Node{
				Key: 3,
			},
			Right: &binaryTree.Node{
				Key: 5,
			},
		},
		Right: &binaryTree.Node{
			Key: 18,
			Left: &binaryTree.Node{
				Key: 12,
			},
			Right: &binaryTree.Node{
				Key: 20,
			},
		},
	}
	FoundNode := binaryTree.Search(tree, 5)
	fmt.Printf("found node containing desired key: %d\n", FoundNode.Key)
	MinimumNode := binaryTree.Minimum(tree)
	fmt.Printf("found minimum node: %d", MinimumNode.Key)
}
