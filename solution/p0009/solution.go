package p0009

import "strconv"

type PalindromeNumberSolution func(x int) bool

func ToStringSolution(palindrom int) bool {
	s := strconv.Itoa(palindrom)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

func DivisionSolution(palindrom int) bool {
	// negative numbers cant be palindrom
	if palindrom < 0 {
		return false
	}
	// numbers that ends with 0 cant be palindrom
	if palindrom%10 == 0 && palindrom != 0 {
		return false
	}
	// single digit numbers are palindrom
	if palindrom < 10 {
		return true
	}

	var reversedHalf int
	for palindrom > reversedHalf {
		reversedHalf = reversedHalf*10 + palindrom%10
		palindrom /= 10
	}

	// second case for numbers with non-even amount of digits
	return palindrom == reversedHalf || palindrom == reversedHalf/10
}
