package p0001

import (
	"fmt"
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
				nums   []int
				target int
				r      []int
			}{
				{[]int{1, 2, 3, 4}, 100, nil}, // No solution
				{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
				{[]int{3, 2, 4}, 6, []int{1, 2}},
				{[]int{3, 3}, 6, []int{0, 1}},
			}

			for _, tc := range tcs {
				t.Run(fmt.Sprintf("%v %d", tc.nums, tc.target), func(t *testing.T) {
					r := solution.fn(tc.nums, tc.target)
					assert.ElementsMatch(t, tc.r, r)
				})
			}
		})
	}
}
