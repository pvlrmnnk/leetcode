package p1137

type NthTribonacciNumberSolution func(n int) int

func Solution(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	arr := make([]int, n+1)
	arr[0], arr[1], arr[2] = 0, 1, 1
	for i := 3; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2] + arr[i-3]
	}

	return arr[n]
}

func OptimizedSolution(n int) int {
	if n < 2 {
		return n
	}

	a, b, c := 0, 1, 1
	for n > 2 {
		a, b, c = b, c, a+b+c
		n--
	}

	return c
}
