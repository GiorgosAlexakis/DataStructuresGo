package binaryTree

func Search(node *Node, key int) *Node {
	if node == nil || key == node.Key {
		return node
	}
	if key < node.Key {
		return Search(node.Left, key)
	} else {
		return Search(node.Right, key)
	}
}

func Minimum(node *Node) *Node {
	for node.Left != nil {
		node = node.Left
	}
	return node
}
