package main

import "fmt"

func main() {
	fmt.Println(isHappy(12345))
}

func isHappy(n int) bool {
	seen := make(map[int]bool)
	for !seen[n] && n != 1 {
		seen[n] = true // Mark the current number as seen
		n = sumSquares(n)
	}
	return n==1
}

func sumSquares(n int) int {
	output := 0
	for n != 0 {
		digit := n % 10
		digit = digit * digit
		output += digit
		n /= 10
	}
	return output
}
