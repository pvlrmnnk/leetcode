package p0009

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindromeNumberSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   PalindromeNumberSolution
	}{
		{"string solution", ToStringSolution},
		{"division solution", DivisionSolution},
	}

	for _, s := range ss {
		s := s
		t.Run(s.name, func(t *testing.T) {
			t.Parallel()

			type args struct {
				x int
			}
			tcs := []struct {
				name string
				args args
				want bool
			}{
				{
					"",
					args{121},
					true,
				},
				{
					"",
					args{-121},
					false,
				},
				{
					"",
					args{10},
					false,
				},
				{
					"",
					args{1},
					true,
				},
			}

			for _, tc := range tcs {
				tc := tc
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					got := s.fn(tc.args.x)
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}
