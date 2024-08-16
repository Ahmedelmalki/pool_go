package main

import "fmt"

func main() {
	// fmt.Println(isIsomorphic("",""))
	fmt.Println(isIsomorphic("ab", "cd"))
	 fmt.Println(isIsomorphic("ufdw","ffuy"))
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mapS := make(map[byte]byte)
	mapT := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if mapS[t[i]] > 0 && mapS[t[i]] != s[i] || mapT[s[i]] > 0 && mapT[s[i]] != t[i] {
			return false
		}
		mapS[t[i]] = s[i]
		mapT[s[i]] = t[i]
	}
	return true
}
