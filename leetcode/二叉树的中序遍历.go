package leetcode

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	stack := list.New()
	res := make([]int, 0)
	for root != nil || stack.Len() != 0 {
		for root != nil {
			stack.PushBack(root)
			root = root.Left
		}
		if stack.Len() != 0 {
			v := stack.Back()
			root = v.Value.(*TreeNode)
			res = append(res, root.Val)
			root = root.Right
			stack.Remove(v)
		}
	}
	return res
}

func newTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}
