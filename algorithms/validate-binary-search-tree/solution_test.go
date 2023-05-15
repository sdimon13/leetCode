package validate_binary_search_tree

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Example 1",
			args{root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			}},
			true,
		},
		{
			"Example 2",
			args{root: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 6},
				},
			}},
			false,
		},
		{
			"Example 3",
			args{root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 2},
			}},
			false,
		},
		{
			"Example 4",
			args{root: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 4},
				Right: &TreeNode{
					Val:   6,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 7},
				},
			}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
