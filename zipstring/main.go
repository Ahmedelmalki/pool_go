package main

import (
	"strconv"
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := []string{
		"aaaa",
		"zzzzzZZZZZZ",
		"",
		"ziiiiipMeee",
		"hhellloTthereYouuunggFelllas",
	}

	for _, arg := range args {
		challenge.Function("ZipString", ZipString, solutions.ZipString, arg)
	}
}

func ZipString(s string) string {
	if len(s) == 0 {
		return s
	}
	rs := ""
	c := 1
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			c++
		} else {
			rs += strconv.Itoa(c) + string(s[i])
			c = 1
		}
	}
	rs += strconv.Itoa(c) + string(s[len(s)-1])
	return rs
}
