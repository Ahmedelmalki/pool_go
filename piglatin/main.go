package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 || len(os.Args) <= 1 {
		fmt.Println()
		return
	}
	args := os.Args[1]
	fmt.Println(piglatinWords(args))
}

func piglatinWords(s string) string {
	rs := ""
	l := len(s)
	if !isthereAvowel(s){
		return "No vowels"
	}
	for i := 0; i < l; i++ {
		if isvowel(rune(s[i])) {
			rs = s[i:l] + s[0:i] + "ay"
			break
		}
	}
	return rs
}

func isvowel(r rune) bool {
	if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' {
		return true
	}
	return false
}

func isthereAvowel(s string) bool {
	for i := 0; i < len(s); i++ {
		if isvowel(rune(s[i])) {
			return true
		}
	}
	return false
}
