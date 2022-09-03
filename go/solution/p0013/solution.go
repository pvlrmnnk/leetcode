package p0013

type RomanToIntSolution func(s string) int

func Solution(roman string) int {
	dic := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var acc int
	runes := []rune(roman)
	for i := 0; i < len(runes); i++ {
		// for last symbol
		if i == len(runes)-1 {
			acc += dic[runes[i]]

			break
		}
		if dic[runes[i]] >= dic[runes[i+1]] {
			acc += dic[runes[i]]
		} else {
			acc -= dic[runes[i]]
		}
	}

	return acc
}

func ReversedOrderSolution(roman string) int {
	dic := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var acc int
	var current, left rune

	for i := len(roman) - 1; i >= 0; i-- {
		current = rune(roman[i])
		if dic[current] < dic[left] {
			acc -= dic[current]
		} else {
			acc += dic[current]
		}
		left = current
	}

	return acc
}
