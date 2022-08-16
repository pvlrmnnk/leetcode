package p0009

import "strconv"

type PalindromeNumberSolution func(x int) bool

func ToStringSolution() PalindromeNumberSolution {
	return func(x int) bool {
		s := strconv.Itoa(x)

		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}

		return true
	}
}

func DivisionSolution() PalindromeNumberSolution {
	return func(x int) bool {
		// negative numbers cant be palindrom
		if x < 0 {
			return false
		}
		// numbers that ends with 0 cant be palindrom
		if x%10 == 0 && x != 0 {
			return false
		}
		// single digit numbers are palindrom
		if x < 10 {
			return true
		}

		var r int
		for x > r {
			r = r*10 + x%10
			x /= 10
		}

		// second case for numbers with non-even amount of digits
		return x == r || x == r/10
	}
}
