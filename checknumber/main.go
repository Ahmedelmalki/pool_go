package main

import "fmt"

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}

func CheckNumber(arg string) bool {
	rs := []rune(arg)
	for i := 0; i < len(rs); i++ {
		if rs[i] == '0' || rs[i] == '1' || rs[i] == '2' || rs[i] == '3' || rs[i] == '3' || rs[i] == '4' || rs[i] == '5' || rs[i] == '6' || rs[i] == '7' || rs[i] == '8' || rs[i] == '9' {
			return true
		}
	}
	return false
}
