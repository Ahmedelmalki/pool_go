package main

import "fmt"

func add(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func main() {
	added := add(2)
	fmt.Println(added(5))
}
