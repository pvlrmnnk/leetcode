package p0121

import (
	"testing"

	"github.com/pvlrmnnk/leetcode-go/common/test"
	"github.com/stretchr/testify/assert"
)

func TestBestTimeToBuyAndSellStockSolution(t *testing.T) {
	ss := []BestTimeToBuyAndSellStockSolution{
		DumbSolution(),
		SmartSolution(),
	}

	for i, s := range ss {
		t.Run(test.N("s", i), func(t *testing.T) {
			tcs := []struct {
				prices []int
				r      int
			}{
				{[]int{7, 1, 5, 3, 6, 4}, 5},
				{[]int{7, 6, 4, 3, 1}, 0},
			}

			for i, tc := range tcs {
				t.Run(test.N("tc", i), func(t *testing.T) {
					r := s(tc.prices)
					assert.Equal(t, tc.r, r)
				})
			}
		})
	}
}
