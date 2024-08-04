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
	ma3lomat := os.Args[1]
	a := os.Args[2]
	b := os.Args[3]
	fmt.Println(searchreplace(ma3lomat, a, b))
}
func searchreplace(s, a, b string) string {
	l := len(s)
	for i := 0; i < len(s); i++ {
		if string(s[i]) == a {
			s = s[0:i] + b + s[i+1:l]
			//fmt.Println(s)
		}
	}
	return s
}
