package main

import (
	"fmt"
	"os"
)

func Firstword(str string) string {
	rs := ""
	i := 0
	for i < len(str) && str[i] == ' ' {
		i++
	}
	for i < len(str) && str[i] != ' ' {
		rs += string(str[i])
		i++
	}
	return rs + "\n"
}

var testCases = []struct {
	in   string
	want string
}{
	{"", "\n"},
	{"             a as", "a\n"},
	{"   f     d", "f\n"},
	{"   asd    ad", "asd\n"},
	{"   salut !!! ", "salut\n"},
	{" salut ! ! !", "salut\n"},
	{"salut ! !", "salut\n"},
}

func main() {
	for _, tc := range testCases {
		got := Firstword(tc.in)
		if got != tc.want {
			fmt.Printf("FirstWord(%q) = %q instead of %q\n", tc.in, got, tc.want)
			os.Exit(1)
		}
	}
}
