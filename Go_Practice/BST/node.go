package main

type Node struct {
	left  *Node
	right *Node
	val   int
}

func (n Node) newNode(val int) *Node {
	return &Node{nil, nil, val}
}