package tree

type BinaryTree struct {
	Root *Node
}

type Node struct {
	Data interface{}
	Left *Node
	Right *Node
}


