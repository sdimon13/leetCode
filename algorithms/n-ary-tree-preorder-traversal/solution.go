package n_ary_tree_preorder_traversal

type Node struct {
	Val      int
	Children []*Node
}

func preorderRecursive(root *Node) []int {
	if root == nil {
		return nil
	}

	result := []int{root.Val}
	for _, node := range root.Children {
		result = append(result, preorderRecursive(node)...)
	}

	return result
}

func preorderIterative(root *Node) []int {
	if root == nil {
		return nil
	}

	stack := []*Node{root}
	result := []int{}

	for len(stack) > 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, root.Val)

		for i := len(root.Children) - 1; i >= 0; i-- {
			stack = append(stack, root.Children[i])
		}
	}
	return result
}
