package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("\n")
	} else {
		// str := os.Args[:1]
		for _, str := range os.Args[1:] {
			fmt.Printf(alphamirror(str))
		}
		fmt.Printf("\n")
	}
}

func alphamirror(s string) string {
	var rs2 []rune
	rs := []rune(s)
	for i := 0; i < len(rs); i++ {
		if isLow(string(rs[i])) {
			rs[i] = ('z' - rs[i]) + 'a'
			rs2 = append(rs2, rs[i])
		}
		if isUp(string(rs[i])) {
			rs[i] = ('Z' - rs[i]) + 'A'
			rs2 = append(rs2, rs[i])
		}
		if !isLow(string(rs[i])) && !isUp(string(rs[i])) {
			rs2 = append(rs2, rs[i])
		}
	}
	return string(rs2)
}

func isLow(s string) bool {
	rs := []rune(s)
	for i := 0; i < len(rs); i++ {
		if rs[i] >= 'a' && rs[i] <= 'z' {
			return true
		}
	}
	return false
}

func isUp(s string) bool {
	rs := []rune(s)
	for i := 0; i < len(rs); i++ {
		if rs[i] >= 'A' && rs[i] <= 'Z' {
			return true
		}
	}
	return false
}
