package main

import "fmt"

func split(s, sep string) []string {
	var words []string
	lastStart := 0
	for i := 0; i < len(s); i++ {
	  if string(s[i]) == sep {
		words = append(words, s[lastStart:i])
		lastStart = i + 1 
	  }
	}
	if lastStart < len(s) {
	  words = append(words, s[lastStart:])
	}
	return words
  }
  

func main() {
	fmt.Println(split("hello world.there more than this,in this", " "))

}
