package main

import (
	"fmt"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := [10]byte{}

	for j := 0; j < 5; j++ {
		for i := 0; i < random.IntBetween(7, 10); i++ {
			table[i] = byte(random.IntBetween(13, 126))
		}
		challenge.Function("PrintMemory", PrintMemory, solutions.PrintMemory, table)
	}
	table2 := [10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'}
	challenge.Function("PrintMemory", PrintMemory, solutions.PrintMemory, table2)
}

func PrintMemory(arr [10]byte) {
	hexline := ""
	asciiline := ""
	for i, b := range arr {
		hexline += toHex(int(b)) + " "
		if b >= 32 && b <= 126 {
			asciiline += string(b)
		} else {
			asciiline += "."
		}
		if (i+1)%4 == 0 || i == len(arr)-1 {
			if len(hexline) > 0 {
				fmt.Println(hexline[:len(hexline)-1])			}
			hexline = ""
		}
	}
	fmt.Println(asciiline)
}

func toHex(n int) string {
	hexDigits := "0123456789abcdef"
	result := ""
	for n > 0 {
		remainder := n % 16
		result = string(hexDigits[remainder]) + result
		n = n / 16
	}
	for len(result) < 2 {
		result = "0" +result
	}
	return result
}
