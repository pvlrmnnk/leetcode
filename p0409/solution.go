package p0409

type LongestPalindromeSolution func(s string) int

func Solution() LongestPalindromeSolution {
	return func(s string) int {
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
		} else {
			return len(s)
		}
	}
}

func GreedySolution() LongestPalindromeSolution {
	return func(s string) int {
		dic := make(map[rune]int)
		for _, r := range s {
			dic[r]++
		}

		l := 0
		for _, c := range dic {
			l += c / 2 * 2
		}

		if l+1 < len(s) {
			return l + 1
		} else {
			return len(s)
		}
	}
}
