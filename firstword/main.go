package main

import "fmt"

func main() {
fmt.Println(firstword("Write a function that takes a string and return a string containing its first word"))
}

func firstword(str string) string {
	words := []string{}
	for i := 0; i < len(str); i++ {
		if str[i] == 32 {
			words = append(words, str[:i])
		}
	}
	return words[0]
}
