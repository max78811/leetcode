package main

// https://leetcode.com/problems/majority-element-ii/
func majorityElement(nums []int) []int {
	m := make(map[int]int) //val count
	for _, i := range nums {
		m[i]++
	}

	result := make([]int, 0, len(m))
	for val, count := range m {
		if count > len(nums)/3 {
			result = append(result, val)
		}
	}
	return result
}
