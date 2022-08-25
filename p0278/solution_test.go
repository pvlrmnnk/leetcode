package p0278

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchSolution(t *testing.T) {
	solutions := []struct {
		name string
		fn   FirstBadVersionSolution
	}{
		{"solution", Solution},
	}

	for _, solution := range solutions {
		t.Run(solution.name, func(t *testing.T) {
			testCases := []struct {
				name       string
				currentVer int
				badVer     int
			}{
				{
					"bad version is 4",
					100, 4,
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					badVer := solution.fn(testCase.currentVer)
					assert.Equal(t, testCase.badVer, badVer)
				})
			}
		})
	}
}
