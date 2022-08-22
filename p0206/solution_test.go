package p0206

import (
	"testing"

	"github.com/pvlrmnnk/leetcode-go/common/test"
	"github.com/stretchr/testify/assert"
)

func TestMergeTwoSortedListsSolution(t *testing.T) {
	ss := []ReverseLinkedListSolution{
		IterativeSolution(),
		RecursiveSolution(),
	}

	for i, s := range ss {
		t.Run(test.N("s", i), func(t *testing.T) {
			tcs := []struct {
				head, r *ListNode
			}{
				{
					&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
					&ListNode{5, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}},
				},
				{
					&ListNode{1, &ListNode{2, nil}},
					&ListNode{2, &ListNode{1, nil}},
				},
				{
					nil, nil,
				},
			}

			for i, tc := range tcs {
				t.Run(test.N("tc", i), func(t *testing.T) {
					r := s(tc.head)
					assert.EqualValues(t, asSlice(t, tc.r), asSlice(t, r))
				})
			}
		})
	}
}

func asSlice(t *testing.T, node *ListNode) []int {
	t.Helper()

	var ints []int
	for node != nil {
		ints = append(ints, node.Val)
		node = node.Next
	}

	return ints
}
