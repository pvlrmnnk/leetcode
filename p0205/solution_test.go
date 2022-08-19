package p0205

import (
	"testing"

	"github.com/pvlrmnnk/leetcode-go/common/test"
	"github.com/stretchr/testify/assert"
)

func TestRunningSumOf1dArraySolution(t *testing.T) {
	ss := []IsomorphicStringsSolution{
		Solution(),
	}

	for i, s := range ss {
		t.Run(test.N("s", i), func(t *testing.T) {
			tcs := []struct {
				s, t string
				r    bool
			}{
				{"egg", "add", true},
				{"foo", "bar", false},
				{"paper", "title", true},
			}

			for i, tc := range tcs {
				t.Run(test.N("tc", i), func(t *testing.T) {
					r := s(tc.s, tc.t)
					assert.Equal(t, tc.r, r)
				})
			}
		})
	}
}
