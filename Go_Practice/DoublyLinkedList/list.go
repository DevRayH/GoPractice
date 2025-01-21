package main

type List struct {
	head *Node
	tail *Node
}

func (l *List) push(val int) {
	var node Node = Node{nil, nil, val}
	if l.head == nil {
		l.head = &node
		l.tail = &node
	} else {
		node.next = l.head
		l.head = &node
	}
}

func (l *List) pop() Node {
	var node Node = *l.head
	l.head = l.head.next
	return node
}