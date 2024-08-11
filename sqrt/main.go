package main

import "fmt"

func main() {
	fmt.Println(Sqrt(81))
	fmt.Println(Sqrt(9))
	fmt.Println(Sqrt(0))
	fmt.Println(Sqrt(10))
}

func Sqrt(nb int) int {
	if nb <= 0 {
		return nb
	}
	i := 1
	for ; i*i <= nb; i++ {
		if i*i == nb {
			break
		}
	}
	return i
}
