package main

import "fmt"

func main() {
	fmt.Println(toHex(1))
}

func toHex(n int) string {
	if n == 0 {
		return "0"
	}

	hexDigits := "0123456789abcdef"
	result := ""

	for n > 0 {
		remainder := n % 16
		n = n / 16
		result = string(hexDigits[remainder]) + result
	}

	return result
}
