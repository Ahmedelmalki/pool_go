package main

import "fmt"


func main() {
	fmt.Println(findNextPrime(1000003000))
	fmt.Println(findNextPrime(1000000111))
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2 ; i*i <= n ; i++ {
		if n%i == 0 {
			return false 
		}
	}
	return true
}

func findNextPrime(n int) int {
	if isPrime(n) {
		return n
	} else {
		for !isPrime(n) {
			n = n +1
		}
	}
	return n
}