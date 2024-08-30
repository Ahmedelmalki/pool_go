package main

import "fmt"

func mapInt(values []int, f func(int) int) []int {
	result := make([]int, len(values))
	for i, v := range values {
		result[i] = f(v)
	}
	return result
}

func main() {
	numbers := []int{1, 2, 3, 4}
	squared := mapInt(numbers, func(x int) int { return x * x })
	fmt.Println(squared)
}
