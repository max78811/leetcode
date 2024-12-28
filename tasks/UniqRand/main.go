package main

import (
	"fmt"
	"math/rand"
)

// Generate list of n uniq elements
func uniqRand(n int) []int {
	m := make(map[int]struct{})
	s := make([]int, 0, n)

	for len(s) <= n {
		randNum := rand.Int()
		_, exists := m[randNum]
		if !exists {
			m[randNum] = struct{}{}
			s = append(s, randNum)
		}
	}

	return s
}

func main() {
	fmt.Println(uniqRand(20))
}
