package main

import "fmt"

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory(arr [10]byte) {
	hexline := ""
	asciiline := ""
	for i, b := range arr {
		hexline += fmt.Sprintf("%02x", b)
		if b >= 32 && b <= 126 {
			asciiline += string(b)
		} else {
			asciiline += "."
		}
		if (i+1)%4 == 0 || i == len(arr)-1 {
			fmt.Println(hexline)
			hexline = ""
		}
	}
	fmt.Println(asciiline)
}
