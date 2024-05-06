package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("plz check ur input agian")
	} else {
		num, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("invalid input")
		} else {
			fmt.Println(ispower2(num))
		}
	}
}

func ispower2(n int) bool {
	if n <= 1 {
		return false
	}
	for n%2 == 0 {
		n /= 2
		if n == 1 {
			return true
		}
	}
	return false
}
