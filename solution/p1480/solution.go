package p1480

type RunningSumOf1dArraySolution func(nums []int) []int

func Solution(nums []int) []int {
	var acc int

	for i, num := range nums {
		acc += num
		nums[i] = acc
	}

	return nums
}
