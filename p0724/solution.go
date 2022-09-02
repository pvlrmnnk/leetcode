package p0724

type PivotIndexSolution func(nums []int) int

func Solution(nums []int) int {
	var acc1, acc2 int

	for _, num := range nums {
		acc1 += num
	}

	for idx, num := range nums {
		acc1 -= num
		if acc1 == acc2 {
			return idx
		}
		acc2 += num
	}

	return -1
}
