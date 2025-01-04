package main

import "sort"

//https://leetcode.com/problems/set-mismatch/description

func findErrorNums(nums []int) []int {
	sort.Ints(nums)
	m := make(map[int]int) //val count
	missedNum := 1         //3

	for _, i := range nums {
		_, exist := m[i]
		if !exist {
			if i == missedNum {
				missedNum++
			}
		}
		m[i]++
	}
	result := make([]int, 0, 2)

	for k, v := range m {
		if v == 2 {
			result = append(result, k)
			break
		}
	}

	result = append(result, missedNum)
	return result
}
