package p0070

type ClimbingStairsSolution func(n int) int

//nolint:varnamelen
func Solution(n int) int {
	a, b := 1, 1

	for n > 0 {
		a, b = b, a+b
		n--
	}

	return a
}
