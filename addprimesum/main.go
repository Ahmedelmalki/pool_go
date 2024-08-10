package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a number")
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	fmt.Println(addprimesum(n))
}

func addprimesum(n int) int {
	sum := 0
	for i := n; i > 1; i-- {
		if isprime(i) {
			sum += i
		}
	}
	return sum
}

func isprime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
