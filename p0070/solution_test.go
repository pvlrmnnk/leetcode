package p0070

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbingStairsSolution(t *testing.T) {
	solutions := []struct {
		name string
		fn   ClimbingStairsSolution
	}{
		{"solution", Solution},
	}

	for _, solution := range solutions {
		t.Run(solution.name, func(t *testing.T) {
			testCases := []struct {
				name string
				n, r int
			}{
				{
					"",
					2, 2,
				},
				{
					"",
					3, 3,
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
