package p0814

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTreePruningSolution func(root *TreeNode) *TreeNode

func Solution(root *TreeNode) *TreeNode {
	var prune func(node *TreeNode) *TreeNode
	prune = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		node.Left = prune(node.Left)
		node.Right = prune(node.Right)

		if node.Val == 0 && node.Left == nil && node.Right == nil {
			node = nil
		}

		return node
	}

	return prune(root)
}
