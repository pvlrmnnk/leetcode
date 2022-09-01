package p0409

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindromeSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   LongestPalindromeSolution
	}{
		{"solution", Solution},
		{"solution", GreedySolution},
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
					args{"abccccdd"},
					7,
				},
				{
					"",
					args{"a"},
					1,
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
