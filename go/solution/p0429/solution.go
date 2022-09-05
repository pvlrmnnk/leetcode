package p0429

import "container/list"

type Node struct {
	Val      int
	Children []*Node
}

type TreeLevelOrderTraversalSolution func(root *Node) [][]int

func Solution(root *Node) [][]int {
	if root == nil {
		return nil
	}

	type wrapper struct {
		node *Node
		lvl  int
	}

	queue := list.New()
	queue.PushBack(&wrapper{root, 0})
	values := make([][]int, 0)

	for queue.Len() != 0 {
		w, _ := queue.Remove(queue.Front()).(*wrapper)
		parent := w.node
		lvl := w.lvl

		if lvl+1 > len(values) {
			values = append(values, make([]int, 0))
		}
		values[lvl] = append(values[lvl], parent.Val)

		for _, child := range parent.Children {
			queue.PushBack(&wrapper{child, lvl + 1})
		}
	}

	return values
}
