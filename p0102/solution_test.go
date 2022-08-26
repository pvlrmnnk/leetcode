package p0102

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTreeLevelOrderTraversalSolution(t *testing.T) {
	solutions := []struct {
		name string
		fn   BinaryTreeLevelOrderTraversalSolution
	}{
		{"solution", Solution},
	}

	for _, solution := range solutions {
		t.Run(solution.name, func(t *testing.T) {
			testCases := []struct {
				name   string
				root   *TreeNode
				values [][]int
			}{
				{
					"",
					nil,
					nil,
				},
				{
					"",
					&TreeNode{1, nil, nil},
					[][]int{{1}},
				},
				{
					"",
					&TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}},
					[][]int{{3}, {9, 20}, {15, 7}},
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					values := solution.fn(testCase.root)
					assert.EqualValues(t, testCase.values, values)
				})
			}
		})
	}
}
