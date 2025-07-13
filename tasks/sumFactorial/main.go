package sumFactorial

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func sumFactorial(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += factorial(i)
	}

	return sum
}
