package main

import (
	"fmt"
	"math/rand"
)

// Generate matrix of uniq elements
func generateMatrix(n, m int) [][]int {
	result := make([][]int, 0, n)
	d := make(map[int]struct{})

	for i := 0; i < n; i++ {
		s := make([]int, 0, m)
		for {
			if len(s) >= m {
				break
			}
			randNum := rand.Int()
			_, exist := d[randNum]
			if exist {
				continue
			}
			d[randNum] = struct{}{}
			s = append(s, randNum)
		}
		result = append(result, s)
	}

	return result
}

func main() {
	fmt.Println(generateMatrix(2, 2))
}
