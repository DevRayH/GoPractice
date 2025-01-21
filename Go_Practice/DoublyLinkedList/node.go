package main

type Node struct {
	next *Node
	last *Node
	val  int
}