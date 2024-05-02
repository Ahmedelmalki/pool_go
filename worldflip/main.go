package main

import (
	"fmt"
)

func main() {
	fmt.Println(WordFlip("First second last"))
	fmt.Println(WordFlip(""))
	fmt.Println(WordFlip("     "))
	fmt.Println(WordFlip(" hello  all  of  you! "))
}

func WordFlip(s string) string {
	words := splitWords(s)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return joinWords(words)
}

func splitWords(s string) []string {
	words := []string{}
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && i != 0 && i != len(s)-1 {
			words = append(words, s[start:i])
			start = i + 1
		}
	}

	if start < len(s) {
		words = append(words, s[start:])
	}
	return words
}

func joinWords(words []string) string {
	result := ""
	for i, word := range words {
		if i > 0 {
			result += " "
		}
		result += word
	}
	return result
}
