package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	wdmatch(os.Args[1], os.Args[2])
}

func wdmatch(s1, s2 string) {
	rs := ""
	for i := 0; i < len(s1); i++ {
		if contains(rune(s1[i]), s2) {
			rs += string(s1[i])
		}
	}
	if s1 == rs {
		fmt.Println(rs)
	}
}

func contains(r rune, s string) bool {
	for i := 0; i < len(s); i++ {
		if r == rune(s[i]) {
			return true
		}
	}
	return false
}
