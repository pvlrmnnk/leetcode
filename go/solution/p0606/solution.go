package p0606

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ConstructStringfromBinaryTreeSolution func(root *TreeNode) string

func Solution(root *TreeNode) string {
	var stringify func(node *TreeNode) string
	stringify = func(node *TreeNode) string {
		switch {
		case node == nil:
			return ""
		case node.Left == nil && node.Right == nil:
			return strconv.Itoa(node.Val)
		case node.Left != nil && node.Right == nil:
			return fmt.Sprintf("%d(%s)", node.Val, stringify(node.Left))
		default:
			return fmt.Sprintf("%d(%s)(%s)", node.Val, stringify(node.Left), stringify(node.Right))
		}
	}

	return stringify(root)
}
