package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type Linkedlist struct {
	head *Node
}

func (list *Linkedlist) insertAtBack(data int) {
	newNode := &Node{data: data, next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (list *Linkedlist) insertFront(data int) {
	if list.head == nil {
		newNode := &Node{data: data, next: nil}
		list.head = newNode
		return
	}

	newNode := &Node{data: data, next: list.head}
	list.head = newNode
}

func (list *Linkedlist) insertAfterValue(afterValue, data int) {
	newNode := &Node{data: data, next: nil}

	current := list.head

	for current != nil {
		if current.data == afterValue {
			newNode.next = current.next
			current.next = newNode
			return
		}
		current = current.next
	}
	fmt.Printf("cannot insert node, after value %d doesn't exist", afterValue)
	fmt.Println()
}

func (list *Linkedlist) insertBeforeValue(beforeValue, data int) {
	if list.head == nil {
		return
	}

	if list.head.data == beforeValue {
		newNode := &Node{data: data, next: list.head}
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		if current.next.data == beforeValue {
			newNode := &Node{data: data, next: current.next}
			current.next = newNode
			return
		}
		current = current.next
	}
	fmt.Printf("cannot insert node, before value %d doesn't exist", beforeValue)
	fmt.Println()
}

func (list *Linkedlist) print() {
	var current *Node = list.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println()
}

func main() {
	myList := Linkedlist{}
	myList.insertFront(10)
	myList.insertFront(20)
	myList.insertFront(30)
	myList.print()
}
