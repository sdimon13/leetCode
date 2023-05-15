package lowest_common_ancestor_of_a_binary_search_tree

import (
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				root: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 0},
						Right: &TreeNode{
							Val:   4,
							Left:  &TreeNode{Val: 3},
							Right: &TreeNode{Val: 5},
						},
					},
					Right: &TreeNode{
						Val:   8,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 9},
					},
				},
				p: &TreeNode{Val: 2},
				q: &TreeNode{Val: 8},
			},
			want: 6,
		},
		{
			name: "Example 2",
			args: args{
				root: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 0},
						Right: &TreeNode{
							Val:   4,
							Left:  &TreeNode{Val: 3},
							Right: &TreeNode{Val: 5},
						},
					},
					Right: &TreeNode{
						Val:   8,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 9},
					},
				},
				p: &TreeNode{Val: 2},
				q: &TreeNode{Val: 4},
			},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{
				root: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 0},
					Right: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 5},
					},
				},
				p: &TreeNode{Val: 2},
				q: &TreeNode{Val: 3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); got.Val != tt.want {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got.Val, tt.want)
			}
		})
	}
}
