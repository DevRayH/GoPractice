package main

import "fmt"

type List struct{
	Head *Node
	//Size int
}

func (l *List) push(val int){
	var node *Node = &Node{val, nil}
	if l.Head == nil {
		l.Head = node
		//l.Size = 1
	} else {
		//l.Size++
		recursiveAdd(l.Head, val)
	}
}

func recursiveAdd(node *Node, val int){
	if node.Next == nil {
		node.Next = newNode(val)
	} else {
		recursiveAdd(node.Next, val)
	}
}

func (l List) walkList(node *Node){
	fmt.Println(node.Val)
	//fmt.Println(l.Size)
	if node.Next != nil{
		l.walkList(node.Next)
	}
}