package p0724

type PivotIndexSolution func(nums []int) int

func Solution() PivotIndexSolution {
	return func(nums []int) int {
		var acc1, acc2 int

		for _, i := range nums {
			acc1 += i
		}

		for idx, i := range nums {
			acc1 -= i
			if acc1 == acc2 {
				return idx
			}
			acc2 += i
		}

		return -1
	}
}
