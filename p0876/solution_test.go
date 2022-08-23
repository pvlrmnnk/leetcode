package p0876

import (
	"testing"

	"github.com/pvlrmnnk/leetcode-go/common/test"
	"github.com/stretchr/testify/assert"
)

func TestMiddleOfTheLinkedListSolution(t *testing.T) {
	ss := []MiddleOfTheLinkedListSolution{
		Solution(),
	}

	for i, s := range ss {
		t.Run(test.N("s", i), func(t *testing.T) {
			tcs := []struct {
				head, r *ListNode
			}{
				func() struct{ head, r *ListNode } {
					head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
					return struct{ head, r *ListNode }{head, head.Next.Next}
				}(),
				func() struct{ head, r *ListNode } {
					head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}
					return struct{ head, r *ListNode }{head, head.Next.Next.Next}
				}(),
				{
					nil, nil,
				},
			}

			for i, tc := range tcs {
				t.Run(test.N("tc", i), func(t *testing.T) {
					r := s(tc.head)
					assert.EqualValues(t, tc.r, r)
				})
			}
		})
	}
}
