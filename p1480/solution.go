package p1480

type RunningSumOf1dArraySolution func(nums []int) []int

func Solution() RunningSumOf1dArraySolution {
	return func(nums []int) []int {
		var acc int

		for i, v := range nums {
			acc += v
			nums[i] = acc
		}

		return nums
	}
}
