package p0338

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingBitsSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   CountingBitsSolution
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
				want []int
			}{
				{
					"",
					args{2},
					[]int{0, 1, 1},
				},
				{
					"",
					args{5},
					[]int{0, 1, 1, 2, 1, 2},
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
