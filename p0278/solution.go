package p0278

type FirstBadVersionSolution func(n int) int

func Solution(v int) int {
	low, high := 1, v

	for low <= high {
		mid := low + (high-low)/2

		if isBadVersion(mid) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}

func isBadVersion(v int) bool {
	return v >= 4
}
