package p0206

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseLinkedListSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   ReverseLinkedListSolution
	}{
		{"iterative solution", IterativeSolution},
		{"recursive solution", RecursiveSolution},
	}

	for _, s := range ss {
		s := s
		t.Run(s.name, func(t *testing.T) {
			t.Parallel()

			type args struct {
				head *ListNode
			}
			tcs := []struct {
				name string
				args args
				want *ListNode
			}{
				{
					"",
					args{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}},
					&ListNode{5, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}},
				},
				{
					"",
					args{&ListNode{1, &ListNode{2, nil}}},
					&ListNode{2, &ListNode{1, nil}},
				},
				{
					"",
					args{nil},
					nil,
				},
			}

			for _, tc := range tcs {
				tc := tc
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					got := s.fn(tc.args.head)
					assert.Equal(t, asSlice(t, tc.want), asSlice(t, got))
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
