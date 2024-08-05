package main

import (
	"fmt"
)

func main() {
	fprime(8333325)
}

func fprime(nums int) {
	i := 2
	for nums > 1 {
		if nums%i == 0 {
			fmt.Print(i, "*")
			nums /= i
		} else {
			i++
		}
	}
}
