package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	wdmatch(os.Args[1], os.Args[2])
}

func wdmatch(s1, s2 string) {
	j := 0
	for i := 0; i < len(s1); i++ {
		kayn := false
		for ; j < len(s2); j++ {
			if s1[i] == s2[j] {
				j++ // move the pointer to make the programm more effective
				kayn = true
				break
			}
		}
		if !kayn {
			return
		}
	}
	fmt.Println(s1)
}
