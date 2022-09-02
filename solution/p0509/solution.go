package p0509

type FibonacciNumberSolution func(n int) int

//nolint:varnamelen
func RecursiveSolution(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return RecursiveSolution(n-1) + RecursiveSolution(n-2)
}

//nolint:varnamelen
func IterativeSolution(n int) int {
	a, b := 0, 1

	for n > 0 {
		a, b = b, a+b
		n--
	}

	return a
}
