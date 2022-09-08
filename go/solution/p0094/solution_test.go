package p0094

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTreeInorderTraversalSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   BinaryTreeInorderTraversalSolution
	}{
		{"", RecursiveSolution},
		{"", IterativeSolution},
		{"", MorrisTraversal},
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
				want []int
			}{
				{
					"",
					args{
						&TreeNode{
							1,
							nil,
							&TreeNode{
								2,
								&TreeNode{3, nil, nil},
								nil,
							},
						},
					},
					[]int{1, 3, 2},
				},
				{
					"",
					args{
						&TreeNode{1, nil, nil},
					},
					[]int{1},
				},
				{
					"",
					args{nil},
					[]int(nil),
				},
			}

			for _, tc := range tcs {
				tc := tc
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					got := s.fn(tc.args.root)
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}
