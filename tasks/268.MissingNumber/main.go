package main

import "sort"

//https://leetcode.com/problems/missing-number/description

func missingNumber(nums []int) int {
	counter := 0
	sort.Ints(nums)

	for _, v := range nums {
		if v != counter {
			return counter
		}
		counter++
	}
	return len(nums)
}
