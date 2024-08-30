package main

import "fmt"

func filter(numbers []int, f func(int) bool) []int {
	var result []int
	for _, value := range numbers {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func iseven(n int) bool {
	return n%2 == 0
}

func main() {
	fmt.Println(filter([]int{1, 2}, iseven))
}
