package p1646

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumInGeneratedArraySolution(t *testing.T) {
	solutions := []struct {
		name string
		fn   MaximumInGeneratedArraySolution
	}{
		{"solution", Solution},
	}

	for _, solution := range solutions {
		t.Run(solution.name, func(t *testing.T) {
			testCases := []struct {
				name string
				n    int
				r    int
			}{
				{
					"",
					7,
					3,
				},
				{
					"",
					2,
					1,
				},
				{
					"",
					3,
					2,
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
