package p0589

import "github.com/emirpasic/gods/stacks/arraystack"

type Node struct {
	Val      int
	Children []*Node
}

type TreePreorderTraversalSolution func(root *Node) []int

func RecursiveSolution(root *Node) []int {
	var walk func(n *Node, v []int) []int

	walk = func(n *Node, v []int) []int {
		if n == nil {
			return v
		}

		v = append(v, n.Val)

		for _, c := range n.Children {
			v = walk(c, v)
		}

		return v
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
