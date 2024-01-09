package bstdepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return maxDepthR(root, 0)
}

func maxDepthR(root *TreeNode, curDepth int) int {
	if root == nil {
		return curDepth
	}

	leftDepth := maxDepthR(root.Left, curDepth+1)
	rightDepth := maxDepthR(root.Right, curDepth+1)

	if leftDepth > rightDepth {
		return leftDepth
	}

	return rightDepth
}
