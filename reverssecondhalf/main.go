package main

import (
	"fmt"
)

func main() {
	fmt.Println(ReverseSecondHalf("This is the 1st half This is the 2nd half"))
	fmt.Println(ReverseSecondHalf(""))
	fmt.Println(ReverseSecondHalf("Hello World"))
}

func ReverseSecondHalf(str string) string {
	Len := len(str)
	if Len == 0 {
		fmt.Println("Invalid Output")
	}
	if Len == 1 {
		return str
	}
	midStr := Len / 2
	str = reverseString(str[midStr:])
	return str
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
