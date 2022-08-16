package p0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSumSolution(t *testing.T) {
	solutions := []struct {
		name string
		fn   TwoSumSolution
	}{
		{"Brute force solution", BruteForceSolution()},
		{"Map solution", MapSolution()},
	}

	for _, solution := range solutions {
		t.Run(solution.name, func(t *testing.T) {
			tcs := []struct {
				name   string
				nums   []int
				target int
				idxs   []int
			}{
				{
					"No solution",
					[]int{1, 2, 3, 4},
					100,
					nil,
				},
				{
					"Leet testcase 1",
					[]int{2, 7, 11, 15},
					9,
					[]int{0, 1},
				},
			}

			for _, tc := range tcs {
				t.Run(tc.name, func(t *testing.T) {
					r := solution.fn(tc.nums, tc.target)
					assert.ElementsMatch(t, tc.idxs, r)
				})
			}
		})
	}
}
