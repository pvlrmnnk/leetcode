package p0704

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchSolution(t *testing.T) {
	solutions := []struct {
		name string
		fn   BinarySearchSolution
	}{
		{"solution", Solution},
		{"generics solution", GenericsSolution[int]},
	}

	for _, solution := range solutions {
		t.Run(solution.name, func(t *testing.T) {
			testCases := []struct {
				name      string
				nums      []int
				target    int
				targetIdx int
			}{
				{
					"target is in nums",
					[]int{-1, 0, 3, 5, 9, 12}, 9, 4,
				},
				{
					"target is not in nums",
					[]int{-1, 0, 3, 5, 9, 12}, 2, -1,
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					targetIdx := solution.fn(testCase.nums, testCase.target)
					assert.Equal(t, testCase.targetIdx, targetIdx)
				})
			}
		})
	}
}
