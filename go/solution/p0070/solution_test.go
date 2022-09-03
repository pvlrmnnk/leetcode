package p0070

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbingStairsSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   ClimbingStairsSolution
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
				want int
			}{
				{
					"",
					args{2},
					2,
				},
				{
					"",
					args{3},
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
