package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

func majorityElement(nums []int) int {
	intmp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		intmp[nums[i]]++
		if intmp[nums[i]] > len(nums)/2 {
			return nums[i]
		}
	}
	return -1
}
