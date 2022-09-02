package p0589

import "github.com/emirpasic/gods/stacks/arraystack"

type Node struct {
	Val      int
	Children []*Node
}

type TreePreorderTraversalSolution func(root *Node) []int

func RecursiveSolution(root *Node) []int {
	var walk func(node *Node, v []int) []int

	walk = func(node *Node, values []int) []int {
		if node == nil {
			return values
		}

		values = append(values, node.Val)

		for _, c := range node.Children {
			values = walk(c, values)
		}

		return values
	}

	return walk(root, nil)
}

func IterativeSolution(root *Node) []int {
	var values []int
	stack := arraystack.New()

	if root != nil {
		stack.Push(root)
	}

	for !stack.Empty() {
		node, _ := stack.Pop()
		values = append(values, node.(*Node).Val)

		for i := len(node.(*Node).Children) - 1; i >= 0; i-- {
			stack.Push(node.(*Node).Children[i])
		}
	}

	return values
}
