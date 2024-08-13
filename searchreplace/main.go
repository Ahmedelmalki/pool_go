package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 || len(os.Args) > 4 {
		fmt.Println()
		return
	}
	s := os.Args[1]
	a := os.Args[2]
	b := os.Args[3]
	fmt.Println(searchreplace(s, a, b))
}
func searchreplace(s, a, b string) string {
	for i := 0; i < len(s); i++ {
		if string(s[i]) == a {
			s = s[0:i] + b + s[i+1:]
		}
	}
	return s
}
