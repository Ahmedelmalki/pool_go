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

func (list *Linkedlist) insertBack(data int) {
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

func (list *Linkedlist) countNodes() (count int) {
	current := list.head
	for current != nil {
		current = current.next
		count++
	}
	return
}

func (list *Linkedlist) findeNodeAt(index int) *Node {
	var count int = 0
	var current *Node = list.head

	for current != nil {
		count++
		current = current.next
	}

	if index <= 0 || index > count {
		return nil
	}

	current = list.head
	for count = 1; count < index; count++ {
		current = current.next
	}
	return current
}

func (list *Linkedlist) deleteFront() {
	if list.head != nil {
		list.head = list.head.next
		fmt.Printf("head node of the list has been deleted\n")
		return
	}
}

func (list *Linkedlist) deleteLast() {
	if list.head == nil {
		fmt.Printf("linked list is empty\n")
	}

	if list.head.next == nil {
		list.head = nil
		fmt.Printf("last node of the linked list has been deleted\n")
		return
	}

	var current *Node = list.head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil
	fmt.Printf("Last node of linked list has been deleted\n")
}

func (list *Linkedlist) deleteBeforeValue(bf int) {
	var previous *Node
	current := list.head

	if current == nil || current.next == nil {
		fmt.Printf("node with before value %d doesn't exist\n", bf)
		return
	}

	for current.next != nil {
		if current.next.data == bf {
			if previous == nil {
				fmt.Printf("node before value %d has data %d and will be deleted\n", bf, current.data)
				list.head = current.next
			} else {
				fmt.Printf("node before value %d has data %d and will be deleted\n", bf, current.data)
				previous.next = current.next
			}
			return
		}
		previous = current
		current = current.next
	}
	fmt.Printf("Node before value %d not found\n", bf)
}

var myList = Linkedlist{}

func main() {
	myList.insertFront(1)
	myList.insertFront(2)
	myList.insertFront(3)
	myList.insertFront(4)
	myList.insertFront(5)
	myList.insertFront(6)
	myList.insertBack(7)
	myList.insertBack(8)
	myList.insertBack(9)
	myList.insertBack(10)
	myList.insertBack(11)
	myList.insertBack(12)
	myList.insertBack(13)
	myList.insertBack(14)
	myList.insertBack(15)
	myList.insertBack(16)
	myList.insertBack(17)
	myList.insertAfterValue(17, 18)
	myList.insertAfterValue(17, 19)
	myList.insertAfterValue(17, 20)
	myList.insertAfterValue(17, 21)
	myList.insertAfterValue(17, 22)
	myList.insertBeforeValue(1, 23)
	myList.insertBeforeValue(1, 24)
	myList.insertBeforeValue(1, 25)
	myList.print()

	// Count the nodes in the linked list
	count := myList.countNodes()
	fmt.Printf("Total number of nodes: %d\n", count)

	// find and print a node at specific index
	indexToFind := 4
	found := myList.findeNodeAt(indexToFind)
	if found != nil {
		fmt.Printf("%d", found.data)
	} else {
		fmt.Printf("node at ndex %d not found \n", indexToFind)
	}
	myList.deleteFront()
	myList.print()
	myList.deleteLast()
	myList.deleteLast()
	myList.deleteLast()
	myList.deleteBeforeValue(2)
	myList.deleteBeforeValue(17)
	myList.deleteBeforeValue(16)
	myList.deleteBeforeValue(1)
	myList.deleteBeforeValue(5)
	myList.print()
}
