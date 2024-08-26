package main

import (
	"fmt"
	"strings"
)

type node struct {
	value int
	next  *node
}

type linkedlist struct {
	head *node
	len  int
}

func (l *linkedlist) add(value int) {
	newNode := new(node)
	newNode.value = value
	if l.head == nil {
		l.head = newNode
	} else {
		iterator := l.head
		for ; iterator != nil; iterator = iterator.next {
		}
		iterator.next = newNode
	}
}

// func (l *linkedlist) remove(value int) {

// }

// func (l *linkedlist) get(value int) *node {

// }

func (l *linkedlist) String() string{
	sb := strings.Builder{}
}

func main() {
	fmt.Println("hello world")

	l := linkedlist{}
	l.add(1)
	fmt.Println(l)
}
