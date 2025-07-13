package prefixSum

// Для списка целых чисел вернуть список того же размера,
// где на i-ой позиции произведение всех элементов списка,
// кроме элемента на позиции i исходного списка.

// [3, 2, 3] -> [2 * 3, 3 * 3, 3 * 2] -> [6, 9, 6]
// [3, 2, 1] -> [2 * 1, 3 * 1, 3 * 2] -> [2, 3, 6]
// [2, 3, 4] -> [3*4, 2*4, 2*3] -> [12, 8, 6]

//     I
// [2, 0, 3] -> [0, 6, 0]
//        I

func solve(n []int) []int {
	r := 1
	zeroCount := 0
	index := 0

	for i, v := range n {
		if v == 0 {
			index = i
			zeroCount++
			continue
		}
		r *= v
	}

	if zeroCount > 1 {
		return make([]int, len(n))
	}
	if zeroCount == 1 {
		result := make([]int, len(n))
		result[index] = r
		return result
	}

	for i, j := range n {
		n[i] = r / j
	}
	return n
}
