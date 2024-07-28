package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	i := os.Args[1]
	fmt.Println(isValidInput(i))
	fmt.Println(shit(i))
}

// func rpncalc(str string) int {
// 	slc := strings.Split(str, " ")

// }

func isValidInput(str string) bool {
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		if !(unicode.IsDigit(r[i]) || (r[i] == ' ' || r[i] == '+' || r[i] == '-' || r[i] == '*' || r[i] == '/' || r[i] == '%')) {
			//fmt.Println(string(r[i]))
			return false
		}
	}
	return true
}

func shit(str string) []int {
	intslc := []int{}
	num1 := 0
	num2 := 0
	if isValidInput(str) {
		for i := 0; i < len(str)-2; i++ {
			if unicode.IsSpace(rune(str[i])) {
				i++
			}
			if unicode.IsDigit(rune(str[i])) && unicode.IsDigit(rune(str[i+1])) {
				num1, _ = strconv.Atoi(string(str[i]))
				num2, _ = strconv.Atoi(string(str[i+1]))
				if string(str[i+2]) == "+" {
					intslc = append(intslc, num1+num2)
				} else if string(str[i+2]) == "-" {
					intslc = append(intslc, num1-num2)
				} else if string(str[i+2]) == "/" {
					intslc = append(intslc, num1/num2)
				} else if string(str[i+2]) == "*" {
					intslc = append(intslc, num1*num2)
				} else if string(str[i+2]) == "%" {
					intslc = append(intslc, num1%num2)
				} else {
					fmt.Println("error")
					os.Exit(0)
				}
			}
		}
	}
	return intslc
}
