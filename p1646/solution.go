package p1646

type MaximumInGeneratedArraySolution func(n int) int

func Solution(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	arr := make([]int, n+1)
	arr[0], arr[1] = 0, 1
	max := arr[1]

	for i := 2; i <= n; i++ {
		// arr[i] = arr[i/2] + arr[i/2+1] * (i%2)
		if i%2 == 0 {
			arr[i] = arr[i/2]
		} else {
			arr[i] = arr[i/2] + arr[i/2+1]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}
