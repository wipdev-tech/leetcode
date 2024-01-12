package sametree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return traverse(p) == traverse(q)
}

func traverse(t *TreeNode) string {
	if t == nil {
		return "n"
	}

	return fmt.Sprint(t.Val) + traverse(t.Left) + traverse(t.Right)
}
