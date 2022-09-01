package p0509

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciNumberSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   FibonacciNumberSolution
	}{
		{"recursive solution", RecursiveSolution},
		{"iterative solution", IterativeSolution},
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
					args{2},
					1,
				},
				{
					"",
					args{3},
					2,
				},
				{
					"",
					args{4},
					3,
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
