package p0121

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBestTimeToBuyAndSellStockSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   BestTimeToBuyAndSellStockSolution
	}{
		{"dumb solution", DumbSolution},
		{"smart solution", SmartSolution},
	}

	for _, s := range ss {
		s := s
		t.Run(s.name, func(t *testing.T) {
			t.Parallel()

			type args struct {
				prices []int
			}
			tcs := []struct {
				name string
				args args
				want int
			}{
				{
					"",
					args{[]int{7, 1, 5, 3, 6, 4}},
					5,
				},
				{
					"",
					args{[]int{7, 6, 4, 3, 1}},
					0,
				},
			}

			for _, tc := range tcs {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					got := s.fn(tc.args.prices)
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}
