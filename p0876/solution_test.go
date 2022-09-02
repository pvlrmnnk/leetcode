package p0876

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiddleOfTheLinkedListSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   MiddleOfTheLinkedListSolution
	}{
		{"solution", Solution},
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
				func() struct {
					name string
					args args
					want *ListNode
				} {
					head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

					return struct {
						name string
						args args
						want *ListNode
					}{
						"",
						args{head},
						head.Next.Next,
					}
				}(),
				func() struct {
					name string
					args args
					want *ListNode
				} {
					head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}

					return struct {
						name string
						args args
						want *ListNode
					}{
						"",
						args{head},
						head.Next.Next.Next,
					}
				}(),
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
					assert.EqualValues(t, tc.want, got)
				})
			}
		})
	}
}
