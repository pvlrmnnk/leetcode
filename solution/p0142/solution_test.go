package p0142

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedListCycleSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   LinkedListCycleSolution
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
					head := &ListNode{3, &ListNode{2, &ListNode{0, &ListNode{-4, nil}}}}
					head.Next.Next.Next.Next = head.Next

					return struct {
						name string
						args args
						want *ListNode
					}{
						"",
						args{head},
						head.Next,
					}
				}(),
				func() struct {
					name string
					args args
					want *ListNode
				} {
					head := &ListNode{1, &ListNode{2, nil}}
					head.Next.Next = head

					return struct {
						name string
						args args
						want *ListNode
					}{
						"",
						args{head},
						head,
					}
				}(),
				{
					"",
					args{&ListNode{1, nil}},
					nil,
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
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}
