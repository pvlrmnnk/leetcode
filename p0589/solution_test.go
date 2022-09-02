package p0589

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreePreorderTraversalSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   TreePreorderTraversalSolution
	}{
		{"recursive solution", RecursiveSolution},
		{"iterative solution", IterativeSolution},
	}

	for _, s := range ss {
		s := s
		t.Run(s.name, func(t *testing.T) {
			t.Parallel()

			type args struct {
				root *Node
			}
			tcs := []struct {
				name string
				args args
				want []int
			}{
				{
					"empty tree",
					args{nil},
					nil,
				},
				{
					"tree case 1",
					args{&Node{1, []*Node{
						{3, []*Node{
							{5, nil},
							{6, nil},
						}},
						{2, nil},
						{4, nil},
					}}},
					[]int{1, 3, 5, 6, 2, 4},
				},
				{
					"tree case 2",
					args{&Node{1, []*Node{
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
							{9, []*Node{
								{13, nil},
							}},
							{10, nil},
						}},
					}}},
					[]int{1, 2, 3, 6, 7, 11, 14, 4, 8, 12, 5, 9, 13, 10},
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
