package main

import (

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	elems := [][]interface{}{
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			1,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			-3,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			2, 4,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			-2, -1,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			2, 0,
		},
	}

	s := random.StrSlice(chars.Words)

	elems = append(elems, []interface{}{s, -len(s) - 10, -len(s) - 5})

	for i := 0; i < 3; i++ {
		s = random.StrSlice(chars.Words)
		elems = append(elems, []interface{}{s, random.IntBetween(-len(s)-10, len(s)+10), random.IntBetween(-len(s)-8, len(s)+10)})
	}

	for _, a := range elems {
		challenge.Function("Slice", Slice, solutions.Slice, a...)
	}
}

func Slice(a []string, nbrs ...int) []string {
	lenght := len(a)
	if len(nbrs) ==0 {
		return nil
	}
	start := nbrs[0]
	end := lenght
	if len(nbrs) == 1 {
		end = lenght
	} else {
		end = nbrs[len(nbrs)-1]
	}
	if start < 0 {
		start = lenght + start
	}
	if end < 0 {
		end = lenght + end 
	}
	if start < 0 {
		start = 0
	}
	if end > lenght{
		end = lenght
	}
	if start > end  {
		return nil
	}
	return a[start:end]
}
