package main

import "fmt"

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat"}))
}

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)
	// looping over every single string making the sortedstr key for strs
	for _, str := range strs {
		sortedstr := string(sort([]byte(str)))
		anagrams[sortedstr] = append(anagrams[sortedstr], str) 
	}
	result := [][]string{}
	for _, group := range anagrams {
		result = append(result, group)
	}
	return result
}

func swap(b []byte, i, j int) {
	temp := b[i]
	b[i] = b[j]
	b[j] = temp
}

func sort(b []byte) []byte {
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b)-1; j++ {
			if b[i] > b[j] {
				swap(b, i, j)
			}
		}
	}
	return b
}
