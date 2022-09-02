package p0142

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedListCycleSolution func(head *ListNode) *ListNode

func Solution(head *ListNode) *ListNode {
	visited := make(map[*ListNode]struct{})

	for head != nil {
		if _, ok := visited[head]; ok {
			return head
		}
		visited[head] = struct{}{}
		head = head.Next
	}

	return nil
}
