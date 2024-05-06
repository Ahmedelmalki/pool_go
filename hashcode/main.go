package main

import (
	"fmt"
)

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}

func HashCode(dec string) string {
	rs := []rune(dec)
	for i := 0; i < len(rs); i++ {
		rs[i] = (rs[i] + rune(len(rs))) % 127
		if rs[i] < 0 {
			rs[i] = rs[i] + 33
		}
	}
	return string(rs)
}
