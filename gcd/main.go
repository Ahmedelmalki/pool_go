package main

import "fmt"

func main() {
	fmt.Println(gcd(42, 10))
	fmt.Println(gcd(42, 12))
	fmt.Println(gcd(14, 77))
	fmt.Println(gcd(17, 3))
}

func gcd(a, b uint) uint {
	for b != 0 {
		r := a % b
		a = b
		b = r
	}
	return a
}
