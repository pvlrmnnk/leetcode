package p1996

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTheNumberofWeakCharactersSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   TheNumberofWeakCharactersSolution
	}{
		{"solution", Solution},
	}

	for _, s := range ss {
		s := s
		t.Run(s.name, func(t *testing.T) {
			t.Parallel()

			type args struct {
				properties [][]int
			}
			tcs := []struct {
				name string
				args args
				want int
			}{
				{
					"",
					args{
						[][]int{{5, 5}, {6, 3}, {3, 6}},
					},
					0,
				},
				{
					"",
					args{
						[][]int{{2, 2}, {3, 3}},
					},
					1,
				},
				{
					"",
					args{
						[][]int{{1, 5}, {10, 4}, {4, 3}},
					},
					1,
				},
				{
					"",
					args{
						[][]int{{7, 7}, {1, 2}, {9, 7}, {7, 3}, {3, 10}, {9, 8}, {8, 10}, {4, 3}, {1, 5}, {1, 5}},
					},
					6,
				},
			}

			for _, tc := range tcs {
				tc := tc
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					got := s.fn(tc.args.properties)
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}
