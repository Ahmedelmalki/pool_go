package main

import "fmt"

func main() {
	fmt.Println(rot14("Hello! How are You?"))
}

func rot14(s string) string {
	rs := []rune(s)
	for i := 0; i < len(rs); i++ {
		rs[i] = shift(rs[i], 14)
	}
	return string(rs)
}

func shift(r rune, shift int) rune {
	if r >= 65 && r <= 90 {
		r = r + rune(shift)
		if r > 90 {
			r = r - rune(26)
		}
	} else if r >= 97 && r <= 122 {
		r = r + rune(shift)
		if r > 122 {
			r = r - rune(26)
		}
	}
	return r
}
