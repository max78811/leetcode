package main

//https://leetcode.com/problems/contains-duplicate/description/

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	for _, v := range m {
		if v > 1 {
			return true
		}
	}
	return false
}
