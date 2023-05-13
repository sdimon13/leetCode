package binary_tree_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, levelSize)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			level[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}

	return result
}
