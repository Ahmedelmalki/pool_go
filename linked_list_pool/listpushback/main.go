package main

import (
	"fmt"
)

func main() {
	l := &List{}

	ListPushBack(l, "how are you")
	ListPushBack(l, "man")
	ListPushBack(l, "Hello")

	temp := l.Head

	for temp != nil {
		fmt.Println(temp.Data)
		fmt.Println()
		temp = temp.Next
	}
	fmt.Println(*l)
	fmt.Println(l.Head)
	fmt.Println(l.Tail)
}

type Node struct {
	Data string
	Next *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func ListPushBack(l *List, data string) {
	newNode := &Node{Data: data}
	if l.Head != nil {
		l.Tail.Next = newNode
		l.Tail = newNode
	} else {
		l.Head = newNode
		l.Tail = newNode
	}
}
