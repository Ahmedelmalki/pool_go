package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
func CamelToSnakeCase(s string) string {
	var rs string
	prevIsLower := false
	if !isCamel(s) {
		return s
	} else {
		for _,r := range s {
			if isLow(string(r)) {
				rs += string(r)
				prevIsLower = true
			} else if isUp(string(r)){
				if prevIsLower {
				rs += "_"
				} 
				rs += string(r)
				prevIsLower = false
			} else {
				rs = rs + string(r)
			}
		}
		return rs
	}
	
}

func isCamel(txt string) bool {
	rs := []rune(txt)
	for i := 0; i < len(rs); i++ {
		if !isLow(string(rs[0])) {
			return false
		} else if isUp(string(rs[i])) {
			if i == 1 || !isLow(string(rs[i-1])) {
				return false
			}
		} else if !isValidCamelCase(rs[i]) {
			return false
		}
	}
	return true
}

func isValidCamelCase(r rune) bool {
	if r == '_' || (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
		return true
	}
	return false
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
