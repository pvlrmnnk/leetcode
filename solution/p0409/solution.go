package p0409

type LongestPalindromeSolution func(s string) int

//nolint:varnamelen
func Solution(s string) int {
	odds := make(map[rune]struct{})
	for _, r := range s {
		if _, ok := odds[r]; ok {
			delete(odds, r)
		} else {
			odds[r] = struct{}{}
		}
	}

	if len(odds) > 0 {
		return len(s) - len(odds) + 1
	}

	return len(s)
}

//nolint:varnamelen
func GreedySolution(s string) int {
	dic := make(map[rune]int)
	for _, r := range s {
		dic[r]++
	}

	possibleLen := 0
	for _, c := range dic {
		possibleLen += c / 2 * 2
	}

	if possibleLen+1 < len(s) {
		return possibleLen + 1
	}

	return len(s)
}
