package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(evalRPN([]string{"123", "1", "+"}))
}

func evalRPN(tokens []string) int {
	var num int
	stack := []int{}
	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			num1 := stack[len(stack)-2]
			num2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			if token == "+" {
				stack = append(stack, num1+num2)
			} else if token == "-" {
				stack = append(stack, num1-num2)
			} else if token == "*" {
				stack = append(stack, num1*num2)
			} else if token == "/" {
				stack = append(stack, num1/num2)
			}
		} else {
			num = Atoi(token)
		}
		stack = append(stack, num)
	}
	return stack[0]
}

func Atoi(s string) int {
	r := []rune(s)
	var rs int
	sign := 1
	i := 0
	for i < len(r) && unicode.IsSpace(r[i]) {
		i++
	}
	if i < len(r) && (r[i] == '-' || r[i] == '+') {
		if r[i] == '-' {
			sign = -1
		}
		i++
	}
	for i < len(r) && unicode.IsDigit(r[i]) {
		rs = rs*10 + int(r[i]-'0')
		if rs*sign > 2147483647 {
			return 2147483647
		} else if rs*sign < -2147483648 {
			return -2147483648
		}
		i++
	}
	return rs * sign
}
