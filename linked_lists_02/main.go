package main

import "fmt"

type node struct {
	next *node
	data int
}

type linkedlist struct {
	head   *node
	lenght int
}

func (l *linkedlist) prepend(value int) {
	newNode := node{data: value}
	if l.head != nil {
		newNode.next = l.head
		l.head = &newNode
		l.lenght++
	} else {
		l.head = &newNode
		l.lenght++
	}
}

func (l *linkedlist) printlist() {
	if l.head == nil {
		fmt.Println("lista khawya")
		return
	}
	currentNode := l.head
	for currentNode != nil {
		fmt.Print(currentNode.data)
		fmt.Print("->")
		currentNode = currentNode.next
	}
	fmt.Println()
}

func (l *linkedlist) deleteValue(value int) {
	if l.head.data == value {
		l.head = l.head.next
		fmt.Println("tmsa7 errass")
		return
	}

	preToDelete := l.head
	for preToDelete.next.data != value {
		if preToDelete.next.next == nil {
			fmt.Println("value makaynach")
			return
		}
		preToDelete = preToDelete.next
	}
	preToDelete.next = preToDelete.next.next
}

func main() {
	mylist := linkedlist{}
	mylist.prepend(1)
	mylist.prepend(2)
	mylist.prepend(3)
	mylist.prepend(4)
	mylist.prepend(5)
	mylist.prepend(6)
	mylist.prepend(7)
	mylist.prepend(8)
	mylist.prepend(9)
	mylist.prepend(10)
	mylist.printlist()
	fmt.Printf("mylist.lenght: %v\n", mylist.lenght)
	valueToDelete := 1
	for i := 0; i < mylist.lenght; i++ {
		mylist.deleteValue(valueToDelete)
		valueToDelete++
	}
	mylist.printlist()
	fmt.Printf("mylist.lenght: %v\n", mylist.lenght)
}
