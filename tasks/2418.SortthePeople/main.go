package main

import "sort"

// ["Mary","John","Emma"]
// [180,165,170]
// [180,170,165]

//https://leetcode.com/problems/sort-the-people/description

func sortPeople(names []string, heights []int) []string {
	m := make(map[int]string) //height name
	for i := 0; i < len(names); i++ {
		m[heights[i]] = names[i]
	}
	sort.Slice(heights, func(i, j int) bool {
		return heights[i] > heights[j]
	})

	result := make([]string, 0, len(names))
	for _, i := range heights {
		result = append(result, m[i])
	}
	return result
}
