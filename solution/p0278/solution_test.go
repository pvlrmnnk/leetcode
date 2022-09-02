package p0278

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   FirstBadVersionSolution
	}{
		{"solution", Solution},
	}

	for _, s := range ss {
		s := s
		t.Run(s.name, func(t *testing.T) {
			t.Parallel()

			type args struct {
				version int
			}
			tcs := []struct {
				name string
				args args
				want int
			}{
				{
					"bad version is 4",
					args{100},
					4,
				},
			}

			for _, tc := range tcs {
				tc := tc
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					got := s.fn(tc.args.version)
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}
