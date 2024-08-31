package main

import "fmt"

func main() {
	fmt.Println(ceanstr("   this is                                a clean string    . "))
}

func ceanstr(s string) string {
	result, slc, temp :="", []string{}, ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) != " " {
			temp += string(s[i])
		} else if temp != "" {
			slc = append(slc, temp)
			temp = ""
		}
	}
	for j := 0 ; j < len(slc) ; j++{
		result += slc[j] + " "
	} 
	return result
}
