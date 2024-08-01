package main

import "fmt"

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	fmt.Println()
}

func PrintNbr(n int) {
	rs := ""
	if n < 0 {
		fmt.Print("-")
		n = -n
	}
	for n != 0 {
		digit := n % 10
		rs = string(digit + 48) + rs
		n /= 10
	}
	for i := len(rs) - 1; i >= 0; i-- {
		fmt.Print(string(rs[i]))
	}
}
