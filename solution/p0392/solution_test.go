package p0392

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequenceSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   IsSubsequenceSolution
	}{
		{"solution", Solution},
	}

	for _, s := range ss {
		s := s
		t.Run(s.name, func(t *testing.T) {
			t.Parallel()

			type args struct {
				s, t string
			}
			tcs := []struct {
				name string
				args args
				want bool
			}{
				{
					"",
					args{"abc", "ahbgdc"},
					true,
				},
				{
					"",
					args{"axc", "ahbgdc"},
					false,
				},
			}

			for _, tc := range tcs {
				tc := tc
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					got := s.fn(tc.args.s, tc.args.t)
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}
