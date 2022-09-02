package p0746

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostClimbingStairsSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   MinCostClimbingStairsSolution
	}{
		{"solution", Solution},
		{"optimized solution", OptimizedSolution},
	}

	for _, s := range ss {
		s := s
		t.Run(s.name, func(t *testing.T) {
			t.Parallel()

			type args struct {
				cost []int
			}
			tcs := []struct {
				name string
				args args
				want int
			}{
				{
					"",
					args{[]int{10, 15, 20}},
					15,
				},
				{
					"",
					args{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}},
					6,
				},
			}

			for _, tc := range tcs {
				tc := tc
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					got := s.fn(tc.args.cost)
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}
