package main

import "fmt"

func main() {
	fmt.Println(ugly(14))
	fmt.Println(ugly(32))
	fmt.Println(ugly(128))
	fmt.Println(ugly(1024))
	fmt.Println(ugly(1337))
}

func ugly(n int) bool {
	rs := []int{}
	i := 2
	for  n > 1 {
		if n%i == 0 {
			rs = append(rs, i)
			n /= i
		} else {
			i++
		}	
	}
	for j := 0; j < len(rs); j++ {
		if rs[j] > 5 {
			return false
		}
	}
	return true
}

