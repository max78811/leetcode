package main

import (
	"fmt"
)

// Дан слайс целых чисел. Напишите функцию remove, удаляющую все нули
//
// Примеры:
// remove([]) -> []
// remove([0]) -> []
// remove([1,0,0,2]) -> [1,2]

func remove(arr []int) []int {
	write := 0
	for _, num := range arr {
		if num != 0 {
			arr[write] = num
			write++
		}
	}

	return arr[:write]
}

func main() {
	fmt.Println(remove([]int{5, 3, 0, 1, 0}))
}
