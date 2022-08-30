package p0338

type CountingBitsSolution func(n int) []int

func Solution(n int) []int {
	counts := make([]int, n+1)
	counts[0] = 0

	for i := 1; i <= n; i++ {
		// i>>1 == i/2
		// i & 1 == i%2
		counts[i] = counts[i>>1] + (i & 1)
	}

	return counts
}
