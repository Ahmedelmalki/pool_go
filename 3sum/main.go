package main

import "fmt"

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4, 9, 8, 7, 6, 5, 4, 3, 2, 1}))
}

func threeSum(n []int) [][]int {
	res := [][]int{}
	n = sort(n)
	for i, a := range n {
		if i > 0 && a == n[i-1] {
			continue
		}
		l, r := i+1, len(n)-1
		for l < r {
			if a+n[l]+n[r] > 0 {
				r -= 1
			} else if a+n[l]+n[r] < 0 {
				l += 1
			} else {
				res = append(res, []int{a, n[l], n[r]})
				l += 1
				for n[l] == n[l-1] && l < r {
					l += 1
				}
			}
		}
	}
	return res
}

func sort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				swap(nums, j, j+1)
			}
		}
	}
	return nums
}

func swap(n []int, i, j int) {
	temp := n[i]
	n[i] = n[j]
	n[j] = temp
}
