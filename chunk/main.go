package main

import (
	"fmt"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func randomSize() []int {
	var randSlice []int
	for i := 0; i <= random.IntBetween(0, 20); i++ {
		randSlice = append(randSlice, random.Int())
	}
	return randSlice
}

func main() {
	type node struct {
		slice []int
		ch    int
	}
	args := []node{
		{
			slice: []int{},
			ch:    0,
		}, {
			slice: []int{1, 2, 3, 4, 5, 6, 7, 8},
			ch:    0,
		},
	}

	for i := 0; i <= 7; i++ {
		value := node{
			slice: randomSize(),
			ch:    random.IntBetween(0, 10),
		}
		args = append(args, value)
	}
	for _, arg := range args {
		challenge.Function("Chunk", Chunk, solutions.Chunk, arg.slice, arg.ch)
	}
}

func Chunk(slice []int, size int) {
	if len(slice) == 0 ||size <= 0 {
		fmt.Println()
		return
	}
	var subslcs [][]int
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		subslcs = append(subslcs, slice[i:end])
	}
	fmt.Println(subslcs)
}
