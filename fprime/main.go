package main

import (
	"fmt"
)

func main() {
	fmt.Println(fprime(8333325))
}

func fprime(nums int) string {
	rs := []string{}
	if nums == 1 {
		return "1"
	}
	for i := 2; i*i <= nums; i++ {
		for nums%i == 0 {
			rs = append(rs, myitoa(i))
			nums /= i
		}
	}
	if nums > 1 {
		rs = append(rs, myitoa(nums))
	}
	return myjion(rs, "*")
}

func myjion(slc []string, sep string) string {
	if len(slc) == 0 {
		return ""
	}
	if len(slc) == 1 {
		return slc[0]
	}
	rs := ""
	for i := 0; i < len(slc)-1; i++ {
		rs = rs + slc[i] + sep
	}
	rs += slc[len(slc)-1]
	return rs
}

func myitoa(n int) string {
	rs := ""
	for n != 0 {
		rs = string(n%10+48) + rs
		n /= 10
	}
	return rs
}
