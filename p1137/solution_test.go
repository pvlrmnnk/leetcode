package p1137

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNthTribonacciNumberSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   NthTribonacciNumberSolution
	}{
		{"solution", Solution},
		{"optimized solution", OptimizedSolution},
	}

	for _, s := range ss {
		s := s
		t.Run(s.name, func(t *testing.T) {
			t.Parallel()

			type args struct {
				n int
			}
			tcs := []struct {
				name string
				args args
				want int
			}{
				{
					"",
					args{4},
					4,
				},
				{
					"",
					args{25},
					1389537,
				},
			}

			for _, tc := range tcs {
				tc := tc
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					got := s.fn(tc.args.n)
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}
