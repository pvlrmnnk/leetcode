package p0409

import (
	"testing"

	"github.com/pvlrmnnk/leetcode-go/common/test"
	"github.com/stretchr/testify/assert"
)

func TestLongestPalindromeSolution(t *testing.T) {
	ss := []LongestPalindromeSolution{
		Solution(),
		GreedySolution(),
	}

	for i, s := range ss {
		t.Run(test.N("s", i), func(t *testing.T) {
			tcs := []struct {
				s string
				r int
			}{
				{"abccccdd", 7},
				{"a", 1},
			}

			for i, tc := range tcs {
				t.Run(test.N("tc", i), func(t *testing.T) {
					r := s(tc.s)
					assert.Equal(t, tc.r, r)
				})
			}
		})
	}
}
