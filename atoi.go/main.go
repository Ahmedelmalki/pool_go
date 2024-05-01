package main

import "fmt"

func main() {
	fmt.Println(atoi("12345"))
	fmt.Println(atoi("0000000012345"))
	fmt.Println(atoi("012 345"))
	fmt.Println(atoi("Hello World!"))
	fmt.Println(atoi("+1234"))
	fmt.Println(atoi("-1234"))
	fmt.Println(atoi("++1234"))
	fmt.Println(atoi("--1234"))
}

func atoi(s string) int {
	rs := []rune(s)
	sing := 1
	var digit int
	if rs[0] == '-' || rs[0] == '+' {
		if rs[0] == '-' {
			sing = -1
		}
		rs = rs[1:]
	}
	for i := 0; i < len(rs); i++ {
		if rs[i] >= '0' && rs[i] <= '9' {
			digit = digit*10 + int(rs[i]-'0')
		} else {
			return 0
		}
	}
	return digit * sing
}
