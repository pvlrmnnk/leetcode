package p0724

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPivotIndexSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   PivotIndexSolution
	}{
		{"solution", Solution},
	}

	for _, s := range ss {
		s := s
		t.Run(s.name, func(t *testing.T) {
			t.Parallel()

			type args struct {
				nums []int
			}
			tcs := []struct {
				name string
				args args
				want int
			}{
				{
					"",
					args{[]int{1, 7, 3, 6, 5, 6}},
					3,
				},
				{
					"",
					args{[]int{1, 2, 3}},
					-1,
				},
				{
					"",
					args{[]int{2, 1, -1}},
					0,
				},
				{
					"",
					args{[]int{-1, -1, -1, 0, 1, 1}},
					0,
				},
			}

			for _, tc := range tcs {
				tc := tc
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					got := s.fn(tc.args.nums)
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}
