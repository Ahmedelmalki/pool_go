package main

import (
	"fmt"
)

func main() {
	fmt.Println(intToRoman(3456))
}

func intToRoman(n int) string {
	if n < 0 || n >= 4000 {
		return ""
	}
	th := []string{"", "M", "MM", "MMM"}
	h := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	t := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	o := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	rs := th[n/1000] + h[n%1000/100] + t[n%100/10] + o[n%10]
	return rs
}
