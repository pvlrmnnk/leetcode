package p0876

type ListNode struct {
	Val  int
	Next *ListNode
}

type MiddleOfTheLinkedListSolution func(head *ListNode) *ListNode

func Solution(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}
