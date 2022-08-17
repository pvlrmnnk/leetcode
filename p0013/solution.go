package p0013

type RomanToIntSolution func(s string) int

func Solution() RomanToIntSolution {
	return func(s string) int {
		m := map[rune]int{
			'I': 1,
			'V': 5,
			'X': 10,
			'L': 50,
			'C': 100,
			'D': 500,
			'M': 1000,
		}

		var acc int
		rs := []rune(s)
		for i := 0; i < len(rs); i++ {
			// for last symbol
			if i == len(rs)-1 {
				acc += m[rs[i]]
				break
			}
			if m[rs[i]] >= m[rs[i+1]] {
				acc += m[rs[i]]
			} else {
				acc -= m[rs[i]]
			}
		}

		return acc
	}
}
