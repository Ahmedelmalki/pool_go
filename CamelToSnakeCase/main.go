package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := []string{
		"CamelCase",
		"camelCase",
		"HelloWorld",
		"132",
		" ",
		"",
		"A",
		"abcs",
		"AbC",
		"AbCEf",
		"abcAree",
		"ahe1Abde",
		"tesTing1",
		"SOME_VARIABLE",
		"ASuperLonGVariableName",
		"thisIsaTestOfCamelCase",
		"aA",
	}
	for _, arg := range args {
		challenge.Function(
			"CamelToSnakeCase",
			CamelToSnakeCase,
			solutions.CamelToSnakeCase,
			arg,
		)
	}
}


func CamelToSnakeCase(s string) string {
	rs := ""
	if anwar(s) && len(s)>0{
		return s
	}
	if len(s) == 0 {
		return ""
	}
	if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z' {
		return s
	}
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			return s
		}
		if i != 0 && (s[i] >= 'A' && s[i] <= 'Z') && (s[i+1] >= 'a' && s[i+1] <= 'z') && i < len(s)-1 && (s[i-1] >= 'a' && s[i-1] <= 'z') {
			rs += "_"
		}
		rs += string(s[i])

	}
	return rs
}

func anwar(s string) bool {
	for i := 0; i <= len(s)-2; i++ {
		if  (s[i] <= 'Z' && s[i] >= 'A') && (s[i+1] <= 'Z' && s[i+1] >= 'A')  {
			return true
		}
	}
	return false
}
