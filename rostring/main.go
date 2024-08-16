package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(rostring("Let there     be light"))
}

func rostring(s string) string {
	s = strings.TrimSpace(s)
	slc := myfields(s) // This automatically splits by any whitespace and removes extra spaces
	if len(slc) == 0 {
		return ""
	}
	rs := append(slc[1:], slc[0])
	return strings.Join(rs, " ")
}

func myfields(s string) []string {
	rs, slc := "", []string{}
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			rs += string(s[i])
		} else if rs != "" {
			slc = append(slc, rs)
			rs = ""
		}
	}
	if rs != "" {
		slc = append(slc, rs)
	}
	return slc
}
