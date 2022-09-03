package p0746

type MinCostClimbingStairsSolution func(cost []int) int

func Solution(cost []int) int {
	cost = append(cost, 0)
	minCosts := make([]int, len(cost))

	minCosts[0] = cost[0]
	minCosts[1] = cost[1]

	for i := 2; i < len(minCosts); i++ {
		var prevStepMinCost int
		if minCosts[i-2] < minCosts[i-1] {
			prevStepMinCost = minCosts[i-2]
		} else {
			prevStepMinCost = minCosts[i-1]
		}
		minCosts[i] = cost[i] + prevStepMinCost
	}

	return minCosts[len(minCosts)-1]
}

func OptimizedSolution(cost []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}

		return b
	}

	first, second := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		current := cost[i] + min(first, second)
		first, second = second, current
	}

	return min(first, second)
}
