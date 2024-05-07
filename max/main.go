package main

import "fmt"

func main() {
	nb := []int{23, 123, 1, 11, 55, 93}
	nb = sorted(nb)
	fmt.Println(Max(nb))
}

func Max(sl []int) int {
	if len(sl) == 0 {
		return 0
	}
	maxv := sl[0]
	for _, v := range sl {
		if v > maxv {
			maxv = v
		}
	}
	return maxv
}

func sorted(n []int) []int {
	for j := 0; j < len(n); j++ {
		for i := 0; i < len(n)-1; i++ {
			if n[i] > n[j] {
				swap(n, i, j)
			}
		}
	}
	return n
}

func swap(n []int, i, j int) {
	temp := n[i]
	n[i] = n[j]
	n[j] = temp
}
