package p0102

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTreeLevelOrderTraversalSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   BinaryTreeLevelOrderTraversalSolution
	}{
		{"solution", Solution},
	}

	for _, s := range ss {
		s := s
		t.Run(s.name, func(t *testing.T) {
			t.Parallel()

			type args struct {
				root *TreeNode
			}
			tcs := []struct {
				name string
				args args
				want [][]int
			}{
				{
					"empty tree",
					args{nil},
					nil,
				},
				{
					"single node tree",
					args{&TreeNode{1, nil, nil}},
					[][]int{{1}},
				},
				{
					"tree case 1",
					args{&TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}},
					[][]int{{3}, {9, 20}, {15, 7}},
				},
				{
					"tree case 2",
					args{&TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, nil}, &TreeNode{3, nil, &TreeNode{5, nil, nil}}}},
					[][]int{{1}, {2, 3}, {4, 5}},
				},
			}

			for _, tc := range tcs {
				tc := tc
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					got := s.fn(tc.args.root)
					assert.EqualValues(t, tc.want, got)
				})
			}
		})
	}
}
