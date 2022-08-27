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
					"empty tree",
					nil,
					nil,
				},
				{
					"single node tree",
					&TreeNode{1, nil, nil},
					[][]int{{1}},
				},
				{
					"tree case 1",
					&TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}},
					[][]int{{3}, {9, 20}, {15, 7}},
				},
				{
					"tree case 2",
					&TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, nil}, &TreeNode{3, nil, &TreeNode{5, nil, nil}}},
					[][]int{{1}, {2, 3}, {4, 5}},
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
