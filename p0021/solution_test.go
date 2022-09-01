package p0021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoSortedListsSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   MergeTwoSortedListsSolution
	}{
		{"solution", Solution},
	}

	for _, s := range ss {
		s := s
		t.Run(s.name, func(t *testing.T) {
			t.Parallel()

			type args struct {
				list1, list2 *ListNode
			}
			tcs := []struct {
				name string
				args args
				want *ListNode
			}{
				{
					"",
					args{
						&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
						&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
					},
					&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}},
				},
				{
					"",
					args{nil, &ListNode{0, nil}},
					&ListNode{0, nil},
				},
				{
					"",
					args{nil, nil},
					nil,
				},
			}

			for _, tc := range tcs {
				tc := tc
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					got := s.fn(tc.args.list1, tc.args.list2)
					assert.EqualValues(t, asSlice(t, tc.want), asSlice(t, got))
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
