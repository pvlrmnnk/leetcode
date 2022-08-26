package p0102

import (
	"github.com/emirpasic/gods/queues/linkedlistqueue"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTreeLevelOrderTraversalSolution func(root *TreeNode) [][]int

func Solution(root *TreeNode) [][]int {
	var values [][]int
	queue := linkedlistqueue.New()

	if root != nil {
		queue.Enqueue(&TreeNode{
			0, root, nil,
		})
	}

	for !queue.Empty() {
		node, _ := queue.Dequeue()
		var levelValues []int
		if left := node.(*TreeNode).Left; left != nil {
			levelValues = append(levelValues, left.Val)
			queue.Enqueue(left)
		}
		if right := node.(*TreeNode).Right; right != nil {
			levelValues = append(levelValues, right.Val)
			queue.Enqueue(right)
		}
		if len(levelValues) > 0 {
			values = append(values, levelValues)
		}
	}

	return values
}
