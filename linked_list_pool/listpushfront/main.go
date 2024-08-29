package main

import "fmt"

func main() {
	l := &List{}

	ListPushFront(l, "Hello")
	ListPushFront(l, "man")
	ListPushFront(l, "how are you")

	temp := l.Head
	for temp != nil {
		fmt.Println(temp.Data)
		temp = temp.Next
	}
}

type Node struct {
	Data any
	Next *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func ListPushFront(l *List, data any) {
	newNode := &Node{Data: data}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode 
	}
}
