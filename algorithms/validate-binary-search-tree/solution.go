package validate_binary_search_tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(node *TreeNode, low int, high int) bool {
	if node == nil {
		return true
	}

	if node.Val <= low || node.Val >= high {
		return false
	}

	return validate(node.Left, low, node.Val) && validate(node.Right, node.Val, high)
}
