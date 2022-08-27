package p0102

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTreeLevelOrderTraversalSolution func(root *TreeNode) [][]int

type Wrapper struct {
	node  *TreeNode
	level int
}

func Solution(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	levelMap := make(map[int][]int)
	queue := list.New()

	if root != nil {
		queue.PushBack(&Wrapper{
			root, 0,
		})
	}

	for queue.Len() != 0 {
		wrapper := queue.Remove(queue.Front()).(*Wrapper)
		node := wrapper.node
		level := wrapper.level
		levelMap[level] = append(levelMap[level], node.Val)
		if left := node.Left; left != nil {
			queue.PushBack(&Wrapper{
				left, level + 1,
			})
		}
		if right := node.Right; right != nil {
			queue.PushBack(&Wrapper{
				right, level + 1,
			})
		}
	}

	values := make([][]int, len(levelMap))
	for k, v := range levelMap {
		values[k] = v
	}

	return values
}
