package main

import "fmt"

func main() {
	slc := []string{}
	hell := Stack(slc) //this line covertes the slice to a stack data structure
	hell.Push("1")
	hell.Push("2")
	hell.Push("3")
	hell.Push("4")
	hell.Push("5")
	for !hell.IsEmpty() {
		 hell.Pop()
	}
	fmt.Println(hell.IsEmpty())
}

type Stack []string

func (s *Stack) IsEmpty() bool {
	if len(*s) == 0 {
		return true
	}
	return false
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

// func (s *Stack) Pop() (string, bool) {
// 	if s.IsEmpty() {
// 		return "", false
// 	} else {
// 		index := len(*s) - 1
// 		element := (*s)[index]
// 		*s = (*s)[:index]
// 		return element, true
// 	}
// }

func (s *Stack) Pop() {
	if !s.IsEmpty() {
		*s = (*s)[:len(*s)-1]
	}
}