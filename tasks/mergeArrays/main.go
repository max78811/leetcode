package main

import "fmt"

func mergeSlices(arrs ...[]int) []int {
	length := 0
	for _, arr := range arrs {
		length += len(arr)
	}
	result := make([]int, 0, length)

	for _, arr := range arrs {
		result = append(result, arr...)
	}
	return result
}

func main() {
	fmt.Println(mergeSlices([]int{1, 2}, []int{3, 4}, []int{5, 6}))
}
