package p0724

import (
	"testing"

	"github.com/pvlrmnnk/leetcode-go/common/test"
	"github.com/stretchr/testify/assert"
)

func TestPivotIndexSolution(t *testing.T) {
	solutions := []PivotIndexSolution{
		Solution(),
	}

	for i, solution := range solutions {
		t.Run(test.Solution(i), func(t *testing.T) {
			tcs := []struct {
				nums []int
				r    int
			}{
				{[]int{1, 7, 3, 6, 5, 6}, 3},
				{[]int{1, 2, 3}, -1},
				{[]int{2, 1, -1}, 0},
				{[]int{-1, -1, -1, 0, 1, 1}, 0},
			}

			for i, tc := range tcs {
				t.Run(test.TestCase(i), func(t *testing.T) {
					r := solution(tc.nums)
					assert.Equal(t, tc.r, r)
				})
			}
		})
	}
}
