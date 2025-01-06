package main

// https://leetcode.com/problems/find-the-difference/description/
func findTheDifference(s string, t string) byte {
	m := make(map[rune]int)
	for _, i := range s {
		m[i]++
	}
	for _, i := range t {
		m[i]++
	}
	for k, v := range m {
		if v%2 != 0 {
			return byte(k)
		}
	}
	return byte(0)
}
