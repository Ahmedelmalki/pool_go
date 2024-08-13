package main

import (
	"fmt"
)

func main() {
	input1 := []uint{2, 3, 1, 1, 4, 0, 0, 0, 1, 9}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}

func CanJump(n []uint) bool {
	if len(n) == 0 {
		return false
	} else if len(n) == 1 {
		return true
	}
	maxreach := uint(0)
	for i := uint(0); i < uint(len(n)); i++ {
		if i > maxreach {
			return false
		}
		if i+n[i] > maxreach {
			maxreach = i + n[i]
		}
		if maxreach >= uint(len(n))-1 {
			return true
		}
	}
	return false
}
