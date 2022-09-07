package p0606

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructStringfromBinaryTreeSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   ConstructStringfromBinaryTreeSolution
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
				want string
			}{
				{
					"",
					args{
						&TreeNode{
							1,
							&TreeNode{
								2,
								&TreeNode{4, nil, nil},
								nil,
							},
							&TreeNode{3, nil, nil},
						},
					},
					"1(2(4))(3)",
				},
				{
					"",
					args{
						&TreeNode{
							1,
							&TreeNode{
								2,
								nil,
								&TreeNode{4, nil, nil},
							},
							&TreeNode{3, nil, nil},
						},
					},
					"1(2()(4))(3)",
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
