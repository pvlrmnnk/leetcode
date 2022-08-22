package p0142

import (
	"testing"

	"github.com/pvlrmnnk/leetcode-go/common/test"
	"github.com/stretchr/testify/assert"
)

func TestMergeTwoSortedListsSolution(t *testing.T) {
	ss := []LinkedListCycleSolution{
		MapSolution(),
	}

	for i, s := range ss {
		t.Run(test.N("s", i), func(t *testing.T) {
			tcs := []struct {
				head, r *ListNode
			}{}

			for i, tc := range tcs {
				t.Run(test.N("tc", i), func(t *testing.T) {
					r := s(tc.head)
					assert.Equal(t, tc.r, r)
				})
			}
		})
	}
}
