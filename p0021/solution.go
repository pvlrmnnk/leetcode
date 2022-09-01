package p0021

type ListNode struct {
	Val  int
	Next *ListNode
}

type MergeTwoSortedListsSolution func(list1 *ListNode, list2 *ListNode) *ListNode

//nolint:gocritic
func Solution(list1, list2 *ListNode) *ListNode {
	nextNode := func() *ListNode {
		var next *ListNode

		if list1 != nil && list2 != nil {
			if list1.Val <= list2.Val {
				next = list1
				list1 = list1.Next
			} else {
				next = list2
				list2 = list2.Next
			}
		} else if list1 == nil && list2 != nil {
			next = list2
			list2 = list2.Next
		} else if list1 != nil && list2 == nil {
			next = list1
			list1 = list1.Next
		}

		return next
	}

	var top, current *ListNode

	current = nextNode()
	top = current
	for current != nil {
		current.Next = nextNode()
		current = current.Next
	}

	return top
}
