package binaryTree

import "log"

func Search(node *Node, key int) *Node {
	if node == nil || key == node.Key {
		return node
	}
	log.Println(node.Key)
	if key < node.Key {
		return Search(node.Left, key)
	} else {
		return Search(node.Right, key)
	}
}
