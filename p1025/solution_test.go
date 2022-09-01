package p1025

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivisorGameSolution(t *testing.T) {
	solutions := []struct {
		name string
		fn   DivisorGameSolution
	}{
		{"solution", Solution},
	}

	for _, solution := range solutions {
		t.Run(solution.name, func(t *testing.T) {
			testCases := []struct {
				name string
				n    int
				r    bool
			}{
				{
					"",
					2, true,
				},
				{
					"",
					3, false,
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
