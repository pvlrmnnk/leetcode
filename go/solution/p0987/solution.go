package p0987

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type VerticalOrderTraversalBinaryTreeSolution func(root *TreeNode) [][]int

func Solution(root *TreeNode) [][]int {
	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}

		return b
	}

	var borders func(node *TreeNode, col int, lvl int) (int, int, int)
	borders = func(node *TreeNode, col int, lvl int) (int, int, int) {
		if node == nil {
			return col + 1, col - 1, lvl - 1
		}
		leftiest, rightiest, depth := col, col, lvl
		leftLeftiest, leftRightiest, leftDepth := borders(node.Left, col-1, lvl+1)
		rightLeftiest, rightRightiest, rightDepth := borders(node.Right, col+1, lvl+1)
		leftiest = min(leftiest, min(leftLeftiest, rightLeftiest))
		rightiest = max(rightiest, max(leftRightiest, rightRightiest))
		depth = max(depth, max(leftDepth, rightDepth))

		return leftiest, rightiest, depth
	}

	leftiest, rightiest, depth := borders(root, 0, 0)
	offset := -1 * leftiest
	width := -1*leftiest + rightiest + 1
	height := depth + 1
	matrix := make([][][]int, width)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([][]int, height)
	}

	var walk func(node *TreeNode, col int, lvl int)
	walk = func(node *TreeNode, col int, lvl int) {
		if node == nil {
			return
		}
		matrix[col+offset][lvl] = append(matrix[col+offset][lvl], node.Val)
		walk(node.Left, col-1, lvl+1)
		walk(node.Right, col+1, lvl+1)
	}

	walk(root, 0, 0)

	vals := make([][]int, len(matrix))
	for i := 0; i < len(vals); i++ {
		lvls := matrix[i]
		for j := 0; j < len(lvls); j++ {
			if len(lvls[j]) == 0 {
				continue
			}
			sort.Ints(lvls[j])
			vals[i] = append(vals[i], lvls[j]...)
		}
	}

	return vals
}
