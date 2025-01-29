package tasks

import "math"

// https://leetcode.com/problems/score-of-a-string/description/
func scoreOfString(s string) int {
	a := []rune(s)
	prev := a[0]
	var result float64

	for i, r := range a {
		if i == 0 {
			continue
		}
		h := prev - r
		result += math.Abs(float64(h))
		prev = r
	}
	return int(result)
}
