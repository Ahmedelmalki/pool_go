package main

import (
	"fmt"
)

func main() {
	fmt.Println(inter("padinton", "paqefwtdjetyiytjneytjoeyjnejeyj"))
	fmt.Println(inter("ddf6vewg64f", "twthgdwthdwfteewhrtag6h4ffdhsd"))
}

func inter(str1, str2 string) string {
	rs := ""
	for i := 0; i < len(str1); i++ {
		if mycontains(str2, string(str1[i])) && !mycontains(rs, string(str1[i])) {
			rs += string(str1[i])
		} else {
			continue
		}
	}
	return rs
}

func mycontains(str1, str2 string) bool {
	if len(str2) == 0 {
		return true
	}
	if len(str1) < len(str2) {
		return false
	}
	for i := 0; i < len(str1); i++ {
		if str1[i:i+len(str2)] == str2 {
			return true
		}
	}
	return false
}
