package p0429

import "container/list"

type Node struct {
	Val      int
	Children []*Node
}

type TreeLevelOrderTraversalSolution func(root *Node) [][]int

func Solution(root *Node) (levels [][]int) {
	if root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() != 0 {
		lvl := make([]int, queue.Len())
		for i := 0; i < len(lvl); i++ {
			node, _ := queue.Remove(queue.Front()).(*Node)
			lvl[i] = node.Val
			for _, child := range node.Children {
				queue.PushBack(child)
			}
		}
		levels = append(levels, lvl)
	}

	return
}
