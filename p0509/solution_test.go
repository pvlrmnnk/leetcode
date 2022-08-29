package p0509

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciNumberSolution(t *testing.T) {
	solutions := []struct {
		name string
		fn   FibonacciNumberSolution
	}{
		{"recursive solution", RecursiveSolution},
		{"iterative solution", IterativeSolution},
	}

	for _, solution := range solutions {
		t.Run(solution.name, func(t *testing.T) {
			testCases := []struct {
				name string
				n, r int
			}{
				{
					"",
					2, 1,
				},
				{
					"",
					3, 2,
				},
				{
					"",
					4, 3,
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
