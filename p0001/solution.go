package p0001

type TwoSumSolution func(nums []int, target int) []int

func BruteForceSolution(nums []int, target int) []int {
	for idx1 := 0; idx1 < len(nums); idx1++ {
		for idx2 := idx1 + 1; idx2 < len(nums); idx2++ {
			if nums[idx1]+nums[idx2] == target {
				return []int{idx1, idx2}
			}
		}
	}

	return nil
}

func MapSolution(nums []int, target int) []int {
	complements := make(map[int]int, len(nums))

	for idx1, num := range nums {
		if idx2, ok := complements[target-num]; ok {
			return []int{idx1, idx2}
		}
		complements[num] = idx1
	}

	return nil
}
