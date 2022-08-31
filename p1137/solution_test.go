package p1137

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNthTribonacciNumberSolution(t *testing.T) {
	solutions := []struct {
		name string
		fn   NthTribonacciNumberSolution
	}{
		{"solution", Solution},
		{"optimized solution", OptimizedSolution},
	}

	for _, solution := range solutions {
		t.Run(solution.name, func(t *testing.T) {
			testCases := []struct {
				name string
				n, r int
			}{
				{
					"",
					4, 4,
				},
				{
					"",
					25, 1389537,
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					r := solution.fn(testCase.n)
					assert.Equal(t, testCase.r, r)
				})
			}
		})
	}
}
