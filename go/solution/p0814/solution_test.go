package p0814

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTreePruningSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   BinaryTreePruningSolution
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
				want *TreeNode
			}{
				{
					"",
					args{
						&TreeNode{
							1,
							nil,
							&TreeNode{
								0,
								&TreeNode{0, nil, nil},
								&TreeNode{1, nil, nil},
							},
						},
					},
					&TreeNode{
						1,
						nil,
						&TreeNode{
							0,
							nil,
							&TreeNode{1, nil, nil},
						},
					},
				},
				{
					"",
					args{
						&TreeNode{
							1,
							&TreeNode{
								0,
								&TreeNode{0, nil, nil},
								&TreeNode{0, nil, nil},
							},
							&TreeNode{
								1,
								&TreeNode{0, nil, nil},
								&TreeNode{1, nil, nil},
							},
						},
					},
					&TreeNode{
						1,
						nil,
						&TreeNode{
							1,
							nil,
							&TreeNode{1, nil, nil},
						},
					},
				},
				{
					"",
					args{
						&TreeNode{
							1,
							&TreeNode{
								1,
								&TreeNode{
									1,
									&TreeNode{0, nil, nil},
									nil,
								},
								&TreeNode{1, nil, nil},
							},
							&TreeNode{
								0,
								&TreeNode{0, nil, nil},
								&TreeNode{1, nil, nil},
							},
						},
					},
					&TreeNode{
						1,
						&TreeNode{
							1,
							&TreeNode{1, nil, nil},
							&TreeNode{1, nil, nil},
						},
						&TreeNode{
							0,
							nil,
							&TreeNode{1, nil, nil},
						},
					},
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
