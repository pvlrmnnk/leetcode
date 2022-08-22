package p0142

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedListCycleSolution func(head *ListNode) *ListNode

func MapSolution() LinkedListCycleSolution {
	return func(head *ListNode) *ListNode {
		m := make(map[*ListNode]struct{})

		for head != nil {
			if _, ok := m[head]; ok {
				return head
			}
			m[head] = struct{}{}
			head = head.Next
		}

		return nil
	}
}
