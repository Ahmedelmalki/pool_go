package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring(" 1234567899999999999"))
}

func lengthOfLongestSubstring(s string) int {
	v := 0
	maxv := 0
	start := 0
	seen := make(map[byte]int)
	for i := 0 ; i <len(s); i++{
		if index, kayn := seen[s[i]] ; kayn && index >= start{
		start = index + 1	
		}
		v = i - start + 1 // updating v if kayn is false
		if v > maxv{
			maxv =v 
		} 
		seen[s[i]] = i // moving to the next index
	}
	return maxv
}
