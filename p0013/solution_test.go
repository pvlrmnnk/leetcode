package p0013

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToIntSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   RomanToIntSolution
	}{
		{"solution", Solution},
		{"reversed order solution", ReversedOrderSolution},
	}

	for _, s := range ss {
		s := s
		t.Run(s.name, func(t *testing.T) {
			t.Parallel()

			type args struct {
				s string
			}
			tcs := []struct {
				name string
				args args
				want int
			}{
				{
					"",
					args{"III"},
					3,
				},
				{
					"",
					args{"LVIII"},
					58,
				},
				{
					"",
					args{"MCMXCIV"},
					1994,
				},
			}

			for _, tc := range tcs {
				tc := tc
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					got := s.fn(tc.args.s)
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}
