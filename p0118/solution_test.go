package p0118

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPascalsTriangleSolution(t *testing.T) {
	solutions := []struct {
		name string
		fn   PascalsTriangleSolution
	}{
		{"solution", Solution},
	}

	for _, solution := range solutions {
		t.Run(solution.name, func(t *testing.T) {
			testCases := []struct {
				name  string
				level int
				r     [][]int
			}{
				{
					"",
					1,
					[][]int{{1}},
				},
				{
					"",
					5,
					[][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					r := solution.fn(testCase.level)
					assert.Equal(t, testCase.r, r)
				})
			}
		})
	}
}
