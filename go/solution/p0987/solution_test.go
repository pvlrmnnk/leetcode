package p0987

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerticalOrderTraversalBinaryTreeSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   VerticalOrderTraversalBinaryTreeSolution
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
					"",
					args{
						&TreeNode{
							3,
							&TreeNode{9, nil, nil},
							&TreeNode{
								20,
								&TreeNode{15, nil, nil},
								&TreeNode{7, nil, nil},
							},
						},
					},
					[][]int{{9}, {3, 15}, {20}, {7}},
				},
				{
					"",
					args{
						&TreeNode{
							1,
							&TreeNode{
								2,
								&TreeNode{4, nil, nil},
								&TreeNode{5, nil, nil},
							},
							&TreeNode{
								3,
								&TreeNode{6, nil, nil},
								&TreeNode{7, nil, nil},
							},
						},
					},
					[][]int{{4}, {2}, {1, 5, 6}, {3}, {7}},
				},
				{
					"",
					args{
						&TreeNode{
							1,
							&TreeNode{
								2,
								&TreeNode{4, nil, nil},
								&TreeNode{6, nil, nil},
							},
							&TreeNode{
								3,
								&TreeNode{5, nil, nil},
								&TreeNode{7, nil, nil},
							},
						},
					},
					[][]int{{4}, {2}, {1, 5, 6}, {3}, {7}},
				},
				{
					"",
					args{
						&TreeNode{
							3,
							&TreeNode{
								1,
								&TreeNode{0, nil, nil},
								&TreeNode{2, nil, nil},
							},
							&TreeNode{
								4,
								&TreeNode{2, nil, nil},
								nil,
							},
						},
					},
					[][]int{{0}, {1}, {3, 2, 2}, {4}},
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
