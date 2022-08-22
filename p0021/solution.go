package p0021

type ListNode struct {
	Val  int
	Next *ListNode
}

type MergeTwoSortedListsSolution func(list1 *ListNode, list2 *ListNode) *ListNode

func Solution() MergeTwoSortedListsSolution {
	return func(list1, list2 *ListNode) *ListNode {
		next := func() *ListNode {
			var nn *ListNode

			if list1 != nil && list2 != nil {
				if list1.Val <= list2.Val {
					nn = list1
					list1 = list1.Next
				} else {
					nn = list2
					list2 = list2.Next
				}
			} else if list1 == nil && list2 != nil {
				nn = list2
				list2 = list2.Next
			} else if list1 != nil && list2 == nil {
				nn = list1
				list1 = list1.Next
			}

			return nn
		}

		var top, cn *ListNode

		cn = next()
		top = cn
		for cn != nil {
			cn.Next = next()
			cn = cn.Next
		}

		return top
	}
}
