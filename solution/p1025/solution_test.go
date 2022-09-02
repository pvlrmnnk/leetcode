package p1025

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivisorGameSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   DivisorGameSolution
	}{
		{"solution", Solution},
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
				want bool
			}{
				{
					"",
					args{2},
					true,
				},
				{
					"",
					args{3},
					false,
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
