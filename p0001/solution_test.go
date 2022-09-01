package p0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSumSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   TwoSumSolution
	}{
		{"brute force solution", BruteForceSolution},
		{"map solution", MapSolution},
	}

	for _, s := range ss {
		s := s
		t.Run(s.name, func(t *testing.T) {
			t.Parallel()

			type args struct {
				nums   []int
				target int
			}
			tcs := []struct {
				name string
				args args
				want []int
			}{
				{
					"no solution",
					args{[]int{1, 2, 3, 4}, 100},
					nil,
				},
				{
					"",
					args{[]int{2, 7, 11, 15}, 9},
					[]int{0, 1},
				},
				{
					"",
					args{[]int{3, 2, 4}, 6},
					[]int{1, 2},
				},
				{
					"",
					args{[]int{3, 3}, 6},
					[]int{0, 1},
				},
			}

			for _, tc := range tcs {
				tc := tc
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					got := s.fn(tc.args.nums, tc.args.target)
					assert.ElementsMatch(t, tc.want, got)
				})
			}
		})
	}
}
