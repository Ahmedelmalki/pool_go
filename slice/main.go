package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)
func main() {
	elems := []struct {
		a    []string
		nbrs []int
	}{
		{[]string{"coding", "algorithm", "ascii", "package", "golang"}, []int{1}},
		{[]string{"coding", "algorithm", "ascii", "package", "golang"}, []int{-3}},
		{[]string{"coding", "algorithm", "ascii", "package", "golang"}, []int{2, 4}},
		{[]string{"coding", "algorithm", "ascii", "package", "golang"}, []int{-2, -1}},
		{[]string{"coding", "algorithm", "ascii", "package", "golang"}, []int{2, 0}},
	}

	// Add random test cases
	for i := 0; i < 3; i++ {
		s := random.StrSlice(chars.Words)
		elems = append(elems, struct {
			a    []string
			nbrs []int
		}{
			s,
			[]int{random.IntBetween(-len(s)-10, len(s)+10), random.IntBetween(-len(s)-8, len(s)+10)},
		})
	}

	for _, e := range elems {
		// Unpack both the slice and the variadic integers for the function call
		args := append([]interface{}{e.a}, interfaceSlice(e.nbrs)...)
		challenge.Function("Slice", Slice, solutions.Slice, args...)
	}
}

// Convert []int to []interface{}
func interfaceSlice(slice []int) []interface{} {
	interSlice := make([]interface{}, len(slice))
	for i, v := range slice {
		interSlice[i] = v
	}
	return interSlice
}


func Slice(a []string, nbrs ...int) []string {
	length := len(a)

	if len(nbrs) == 0 {
		return nil
	}

	start := nbrs[0]
	end := length

	if len(nbrs) > 1 {
		end = nbrs[1]
	}

	// Handle negative indices
	if start < 0 {
		start = length + start
	}
	if end < 0 {
		end = length + end
	}

	// Ensure start and end are within bounds
	if start < 0 {
		start = 0
	}
	if end > length {
		end = length
	}

	// If start is greater than end, return nil
	if start > end {
		return nil
	}

	return a[start:end]
}
