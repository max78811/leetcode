package main

func majorityElement(nums []int) int {
	m := make(map[int]int) //val count
	for _, i := range nums {
		m[i]++
	}

	n := len(nums) / 2
	major := 0
	for v, count := range m {
		if count > n {
			major = v
		}
	}

	return major
}
