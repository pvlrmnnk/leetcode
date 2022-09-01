package p0119

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPascalsTriangleSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   PascalsTriangleSolution
	}{
		{"solution", Solution},
	}

	for _, s := range ss {
		s := s
		t.Run(s.name, func(t *testing.T) {
			t.Parallel()

			type args struct {
				level int
			}
			tcs := []struct {
				name string
				args args
				want []int
			}{
				{
					"",
					args{0},
					[]int{1},
				},
				{
					"",
					args{1},
					[]int{1, 1},
				},
				{
					"",
					args{3},
					[]int{1, 3, 3, 1},
				},
			}

			for _, tc := range tcs {
				tc := tc
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					got := s.fn(tc.args.level)
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}
