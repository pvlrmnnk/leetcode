package p0206

type ListNode struct {
	Val  int
	Next *ListNode
}

type ReverseLinkedListSolution func(head *ListNode) *ListNode

func IterativeSolution() ReverseLinkedListSolution {
	return func(head *ListNode) *ListNode {
		var r, next *ListNode

		for head != nil {
			next = head.Next
			head.Next = r
			r = head
			head = next
		}

		return r
	}
}

func RecursiveSolution() ReverseLinkedListSolution {
	var r func(h *ListNode, nh *ListNode) *ListNode

	r = func(h *ListNode, nh *ListNode) *ListNode {
		if h == nil {
			return nh
		}

		next := h.Next
		h.Next = nh

		return r(next, h)
	}

	return func(head *ListNode) *ListNode {
		return r(head, nil)
	}
}
