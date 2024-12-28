package main

//https://leetcode.com/problems/single-number

func singleNumber(nums []int) int {
	m := make(map[int]int) // val count
	for _, i := range nums {
		m[i]++
	}

	for k, count := range m {
		if count == 1 {
			return k
		}
	}
	return -1
}
