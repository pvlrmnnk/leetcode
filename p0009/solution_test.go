package p0009

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindromeNumberSolution(t *testing.T) {
	solutions := []struct {
		name string
		fn   PalindromeNumberSolution
	}{
		{"ToStringSolution", ToStringSolution()},
		{"DivisionSolution", DivisionSolution()},
	}

	for _, solution := range solutions {
		t.Run(solution.name, func(t *testing.T) {
			tcs := []struct {
				x int
				r bool
			}{
				{121, true},
				{-121, false},
				{10, false},
				{1, true},
			}

			for _, tc := range tcs {
				t.Run(fmt.Sprintf("%d", tc.x), func(t *testing.T) {
					r := solution.fn(tc.x)
					assert.Equal(t, tc.r, r)
				})
			}
		})
	}
}
