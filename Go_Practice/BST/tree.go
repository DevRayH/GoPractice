package main

import "fmt"

type Tree struct {
	root *Node
}

func (t *Tree) addNode(val int) {
	var node Node = Node{}
	if t.root == nil {
		t.root = node.newNode((val))
	} else {
		recursiveAdd(t.root, val)
	}
}

func recursiveAdd(node *Node, val int) {
	if node.val > val {
		if node.left == nil {
			node.left = node.newNode(val)
		} else {
			recursiveAdd(node.left, val)
		}
	} else if node.val < val {
		if node.right == nil {
			node.right = node.newNode(val)
		} else {
			recursiveAdd(node.right, val)
		}
	} else {
		var thisNode = node.newNode(val)
		thisNode.val = val
		thisNode.left = node.left
		node.left = thisNode
	}
}
func (t Tree) inorder(){
	recursiveInorder(t.root)
}

func recursiveInorder(node *Node) {
	if node == nil{
		return
	}
	if node.left != nil {
		recursiveInorder(node.left)
	}
	fmt.Println(node.val)

	if node.right != nil{
		recursiveInorder(node.right)
	}
}