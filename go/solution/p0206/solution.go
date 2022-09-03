package p0206

type ListNode struct {
	Val  int
	Next *ListNode
}

type ReverseLinkedListSolution func(head *ListNode) *ListNode

func IterativeSolution(head *ListNode) *ListNode {
	var root, next *ListNode

	for head != nil {
		next = head.Next
		head.Next = root
		root = head
		head = next
	}

	return root
}

func RecursiveSolution(head *ListNode) *ListNode {
	var reverse func(head, next *ListNode) *ListNode

	reverse = func(head *ListNode, next *ListNode) *ListNode {
		if head == nil {
			return next
		}

		next, head.Next = head.Next, next

		return reverse(next, head)
	}

	return reverse(head, nil)
}
