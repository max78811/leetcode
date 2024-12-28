package main

import "sort"

//https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func removeDuplicates(nums []int) int {
	uniq := make(map[int]struct{})
	for _, num := range nums {
		uniq[num] = struct{}{}
	}

	s := make([]int, 0, len(uniq))
	for i, _ := range uniq {
		s = append(s, i)
	}
	sort.Ints(s)

	var count int
	for _, j := range s {
		nums[count] = j
		count++
	}
	return len(uniq)
}
