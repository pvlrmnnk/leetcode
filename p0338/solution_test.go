package p0338

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingBitsSolution(t *testing.T) {
	solutions := []struct {
		name string
		fn   CountingBitsSolution
	}{
		{"solution", Solution},
	}

	for _, solution := range solutions {
		t.Run(solution.name, func(t *testing.T) {
			testCases := []struct {
				name string
				n    int
				r    []int
			}{
				{
					"",
					2,
					[]int{0, 1, 1},
				},
				{
					"",
					5,
					[]int{0, 1, 1, 2, 1, 2},
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
