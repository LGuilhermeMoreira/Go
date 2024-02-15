package tree

import (
	"fmt"
)

type Node struct {
	Value int
	Right *Node
	Left  *Node
}

type binaryTree struct {
	Root *Node
}

func Start() *binaryTree {
	return &binaryTree{
		Root: nil,
	}
}

func addRecursive(root *Node, value int) *Node {

	if root == nil {
		return &Node{
			Value: value,
			Right: nil,
			Left:  nil,
		}

	} else {
		if value >= root.Value {
			root.Right = addRecursive(root.Right, value)
		} else {
			root.Left = addRecursive(root.Left, value)
		}
	}

	return root
}

func (b *binaryTree) Add(value int) {
	b.Root = addRecursive(b.Root, value)
}

func showRecursive(root *Node) {
	if root != nil {
		fmt.Println(root.Value)
		showRecursive(root.Left)
		showRecursive(root.Right)
	}
}

func (b binaryTree) Show() {
	showRecursive(b.Root)
}
