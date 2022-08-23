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
			}{
				func() struct{ head, r *ListNode } {
					head := &ListNode{3, &ListNode{2, &ListNode{0, &ListNode{-4, nil}}}}
					head.Next.Next.Next.Next = head.Next
					return struct{ head, r *ListNode }{head, head.Next}
				}(),
				func() struct{ head, r *ListNode } {
					head := &ListNode{1, &ListNode{2, nil}}
					head.Next.Next = head
					return struct{ head, r *ListNode }{head, head}
				}(),
				{
					&ListNode{1, nil}, nil,
				},
				{
					nil, nil,
				},
			}

			for i, tc := range tcs {
				t.Run(test.N("tc", i), func(t *testing.T) {
					r := s(tc.head)
					assert.Equal(t, tc.r, r)
				})
			}
		})
	}
}
