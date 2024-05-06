package main

import (
	"fmt"
)

func main() {
	fmt.Println(ByeByeFirst([]string{}))
	fmt.Println(ByeByeFirst([]string{"one arg"}))
	fmt.Println(ByeByeFirst([]string{"first", "second"}))
	fmt.Println(ByeByeFirst([]string{"", "abcd", "efg"}))
}

func ByeByeFirst(strings []string) []string {
	if len(strings) == 0 {
		return []string{}
	} else {
		return strings[1:]
	}
}
