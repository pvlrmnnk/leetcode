package p0094

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTreeInorderTraversalSolution func(root *TreeNode) []int

// Time complexity: O(n)
// Space complexity: O(n)
func RecursiveSolution(root *TreeNode) []int {
	var vals []int

	var walk func(node *TreeNode)
	walk = func(node *TreeNode) {
		if node == nil {
			return
		}

		walk(node.Left)
		vals = append(vals, node.Val)
		walk(node.Right)
	}

	walk(root)

	return vals
}

// Time complexity: O(n)
// Space complexity: O(n)
func IterativeSolution(root *TreeNode) []int {
	var vals []int

	stack := list.New()
	node := root
	for node != nil || stack.Len() != 0 {
		for node != nil {
			stack.PushBack(node)
			node = node.Left
		}

		node, _ = stack.Remove(stack.Back()).(*TreeNode)
		vals = append(vals, node.Val)

		node = node.Right
	}

	return vals
}

// Time complexity: O(n)
// Space complexity: O(1)
func MorrisTraversal(root *TreeNode) []int {
	var vals []int
	var current, previous *TreeNode

	current = root
	for current != nil {
		if current.Left == nil {
			vals = append(vals, current.Val)
			current = current.Right
		} else {
			previous = current.Left
			for previous.Right != nil && previous.Right != current {
				previous = previous.Right
			}

			if previous.Right == nil {
				previous.Right = current
				current = current.Left
			} else {
				previous.Right = nil
				vals = append(vals, current.Val)
				current = current.Right
			}
		}
	}

	return vals
}
