package main

import (
	"fmt"
	"strconv"
	"strings"
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
			rs = append(rs ,strconv.Itoa(i))
			nums /= i
		} 	
	}
	if nums > 1 {
		rs = append(rs, strconv.Itoa(nums))
	}
	return strings.Join(rs, "*")
}
