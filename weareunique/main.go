package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := [][]string{
		{"abc", "def"},
		{"hello", "yoall"},
		{"everyone", ""},
		{"hello world", "fam"},
		{"abc", "abc"},
		{"", ""},
		{"pomme", "pomme"},
		{"+265", "265"},
		{"123231", "123231"},
		{"w^p@@j", "w^p@@j"},
		{"26235e5", "4478q92"},
		{"		", "		 "},
		{"AB$%d.52", "eepqdl.52"},
		{"", "eveRyone"},
		{"_55w1se", "55w1se"},
	}
	for _, arg := range table {
		challenge.Function("WeAreUnique", WeAreUnique, solutions.WeAreUnique, arg[0], arg[1])
	}
}

func WeAreUnique(str1, str2 string) int {
	result := ""
	if len(str1) == 0 && len(str2) == 0 {
		return -1
	}
	newstr := str1 + str2
	for i, j := 0, 1; i < len(newstr) && j < len(newstr); i, j = i+1, j+1 {
		if newstr[i] != newstr[j] {
			result += string(newstr[i])
		}
	}
	return len(result)
}

// func mycontains(str1, str2 string) bool {
// 	if len(str2) == 0 {
// 		return true
// 	}
// 	if len(str1) < len(str2) {
// 		return false
// 	}
// 	for i := 0; i < len(str1); i++ {
// 		if str2 == str1[i:i+len(str2)] {
// 			return true
// 		}
// 	}
// 	return false
// }
