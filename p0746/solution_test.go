package p0746

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostClimbingStairsSolution(t *testing.T) {
	solutions := []struct {
		name string
		fn   MinCostClimbingStairsSolution
	}{
		{"solution", Solution},
		{"optimized solution", OptimizedSolution},
	}

	for _, solution := range solutions {
		t.Run(solution.name, func(t *testing.T) {
			testCases := []struct {
				name string
				cost []int
				r    int
			}{
				{
					"",
					[]int{10, 15, 20},
					15,
				},
				{
					"",
					[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
					6,
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					r := solution.fn(testCase.cost)
					assert.Equal(t, testCase.r, r)
				})
			}
		})
	}
}
