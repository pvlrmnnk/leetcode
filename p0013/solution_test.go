package p0013

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToIntSolution(t *testing.T) {
	solutions := []struct {
		name string
		fn   RomanToIntSolution
	}{
		{"Solution", Solution()},
		{"ReversedOrderSolution", ReversedOrderSolution()},
	}

	for _, solution := range solutions {
		t.Run(solution.name, func(t *testing.T) {
			tcs := []struct {
				s string
				r int
			}{
				{"III", 3},
				{"LVIII", 58},
				{"MCMXCIV", 1994},
			}

			for _, tc := range tcs {
				t.Run(fmt.Sprintf("%s", tc.s), func(t *testing.T) {
					r := solution.fn(tc.s)
					assert.Equal(t, tc.r, r)
				})
			}
		})
	}
}
