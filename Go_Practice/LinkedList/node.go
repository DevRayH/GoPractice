package main

type Node struct {
	Val  int
	Next *Node
}

func newNode(val int) *Node {
	return &Node{val, nil}
}