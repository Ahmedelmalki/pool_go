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
	if len(str1) == 0 && len(str2) == 0 {
		return -1
	}
	u1 := uniquechars(str1, str2)
	u2 := uniquechars(str2, str1)
	return len(u1) + len(u2)
}

func uniquechars(str1, str2 string) string {
	uniquechars := []rune{}
	for i := 0; i < len(str1); i++ {
		if !contains(str2, rune(str1[i])) && !containschar(uniquechars, rune(str1[i])) {
			uniquechars = append(uniquechars, rune(str1[i]))
		}
	}
	return string(uniquechars)
}

func contains(s string, r rune) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}
	return false
}

func containschar(slice []rune, r rune) bool {
	for _, v := range slice {
		if v == r {
			return true
		}
	}
	return false
}
