package p0589

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreePreorderTraversalSolution(t *testing.T) {
	solutions := []struct {
		name string
		fn   TreePreorderTraversalSolution
	}{
		{"recursive solution", RecursiveSolution},
		{"iterative solution", IterativeSolution},
	}

	for _, solution := range solutions {
		t.Run(solution.name, func(t *testing.T) {
			testCases := []struct {
				name   string
				root   *Node
				result []int
			}{
				{
					"empty tree",
					nil, nil,
				},
				{
					"tree case 1",
					&Node{1, []*Node{
						{3, []*Node{
							{5, nil},
							{6, nil},
						}},
						{2, nil},
						{4, nil},
					}},
					[]int{1, 3, 5, 6, 2, 4},
				},
				{
					"tree case 2",
					&Node{1, []*Node{
						{2, nil},
						{3, []*Node{
							{6, nil},
							{7, []*Node{
								{11, []*Node{
									{14, nil},
								}},
							}},
						}},
						{4, []*Node{
							{8, []*Node{
								{12, nil},
							}},
						}},
						{5, []*Node{
							{9, *&[]*Node{
								{13, nil},
							}},
							{10, nil},
						}},
					},
					},
					[]int{1, 2, 3, 6, 7, 11, 14, 4, 8, 12, 5, 9, 13, 10},
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					result := solution.fn(testCase.root)
					assert.EqualValues(t, testCase.result, result)
				})
			}
		})
	}
}
