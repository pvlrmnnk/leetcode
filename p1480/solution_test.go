package p1480

import (
	"testing"

	"github.com/pvlrmnnk/leetcode-go/common/test"
	"github.com/stretchr/testify/assert"
)

func TestRunningSumOf1dArraySolution(t *testing.T) {
	solutions := []RunningSumOf1dArraySolution{
		Solution(),
	}

	for i, solution := range solutions {
		t.Run(test.Solution(i), func(t *testing.T) {
			tcs := []struct {
				nums []int
				r    []int
			}{
				{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
				{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
				{[]int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
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
